package tool

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/invopop/jsonschema"
	"github.com/sashabaranov/go-openai"
	"github.com/tealeg/xlsx"
)

func processXLSX(ctx context.Context, apiKey, model, filePath, prompt string, page, firstRow, LastRow, maxTokens int) error {
	log.Printf("process xlsx file: %s, first: %d, last: %d", filePath, firstRow, LastRow)
	client := openai.NewClient(apiKey)

	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		return err
	}

	tasks, err := getGolangTask(ctx, client, model, prompt, LastRow-firstRow+1, maxTokens)
	if err != nil {
		return err
	}

	// repeat the getGolangTask if the number of tasks is not equal to the number of rows
	for len(tasks) < LastRow-firstRow+1 {
		newTasks, err := getGolangTask(ctx, client, model, prompt, LastRow-firstRow+1, maxTokens)
		if err != nil {
			return err
		}
		tasks = append(tasks, newTasks...)
	}

	sheet := file.Sheets[page-1] // page 2
	j := 0
	for i := firstRow; i <= LastRow && j < len(tasks); i++ {
		row := sheet.Rows[i]

		row.Cells[4].SetString(tasks[j])
		j++
	}
	fmt.Printf("write %d tasks\n", j)

	return file.Save(filePath)
}

// FunctionGenerateTasks is a function that generates a task for a developer.
// it takes a task as a parameter and returns a task
type FunctionGenerateTasks struct {
	Tasks []string `json:"tasks" jsonschema:"required,description=golang developer tasks"`
}

func getGolangTask(ctx context.Context, client *openai.Client, model, prompt string, nTaches, maxTokens int) ([]string, error) {
	fmt.Printf("generate %d tasks\n", nTaches)
	promptRoleSystem := fmt.Sprintf(prompt, nTaches)

	req := openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: promptRoleSystem,
			},
		},
		MaxTokens: maxTokens,
		Tools: []openai.Tool{
			{
				Type: openai.ToolTypeFunction,
				Function: openai.FunctionDefinition{
					Name:       "generate_tasks",
					Parameters: jsonschema.Reflect(FunctionGenerateTasks{}).Definitions["FunctionGenerateTasks"],
				},
			},
		},
	}
	if maxTokens > 0 {
		req.MaxTokens = maxTokens
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cannot create chat: %w", err)
	}

	if len(resp.Choices) > 0 {
		if len(resp.Choices[0].Message.ToolCalls) > 0 {
			args := resp.Choices[0].Message.ToolCalls[0].Function.Arguments
			fmt.Println("args: ", args)
			var v FunctionGenerateTasks
			err := json.Unmarshal([]byte(args), &v)
			if err != nil {
				return nil, fmt.Errorf("cannot unmarshal generate-task args: %w", err)
			}

			return v.Tasks, nil
		} else {
			return nil, fmt.Errorf("got content: %s", resp.Choices[0].Message.Content)
		}
	}

	return nil, fmt.Errorf("aucune réponse obtenue de l'API OpenAI: %s", resp.Choices[0])
}

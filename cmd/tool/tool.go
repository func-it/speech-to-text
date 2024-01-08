package tool

import (
	"context"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"

	"github.com/func-it/speechToText/config"
	"github.com/func-it/speechToText/pkg/pretty"
)

func NewToolCmd(ctx context.Context, conf *config.Conf) *cobra.Command {
	cmd := &cobra.Command{
		Use: "tool",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().PersistentPreRunE(cmd.Parent(), args)
		},
	}

	cmd.AddCommand([]*cobra.Command{
		NewAppendXSLXCommand(ctx, conf),
		NewGenLogoCommand(ctx, conf),
		{
			Use:   "config-print",
			Short: "print config",
			RunE: func(cmd *cobra.Command, args []string) error {
				pretty.PrintYaml(conf)
				return nil
			},
		},
		{
			Use:   "modify-file [file-path] [receiver] [function-name] [new-content]",
			Short: "modify file",
			Args:  cobra.ExactArgs(4),
			RunE: func(cmd *cobra.Command, args []string) error {
				return modifyMethodInFile(args[0], args[1], args[2], args[3])
			},
		},
	}...)

	return cmd
}

// NewGenLogoCommand returns a new logo command.
func NewGenLogoCommand(ctx context.Context, conf *config.Conf) *cobra.Command {
	var outPath string
	cmd := &cobra.Command{
		Use:   "gen-logo",
		Short: "print logo",
		RunE: func(cmd *cobra.Command, args []string) error {
			genLogo(outPath)
			return nil
		},
	}

	// bind out path flag to command
	cmd.Flags().StringVar(&outPath, "out-path", "out_gen_img.png", "out path")

	return cmd
}

// NewAppendXSLXCommand returns a new append xslx command.
func NewAppendXSLXCommand(ctx context.Context, conf *config.Conf) *cobra.Command {
	var model string
	var filePath, prompt string
	var firstRow, lastRow, maxTokens, page int

	cmd := &cobra.Command{
		Use:   "append-xslx [gpt-model] [file-path]",
		Short: "append xslx file",
		//Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return processXLSX(ctx, conf.OpenAi.ApiKey, model, filePath, prompt, page, firstRow, lastRow, maxTokens)
		},
	}

	// bind page flag to command
	cmd.Flags().IntVar(&page, "page", 0, "page")

	// bind model flag to command, default is openai.GPT4
	cmd.Flags().StringVar(&model, "model", openai.GPT4TurboPreview, "gpt model")

	// bind file path flag to command
	cmd.Flags().StringVar(&filePath, "file-path", "", "file path")

	// bind first row flag to command
	cmd.Flags().IntVar(&firstRow, "first-row", 0, "first row")

	// bind last row flag to command
	cmd.Flags().IntVar(&lastRow, "last-row", 0, "last row")

	// bind max tokens flag to command
	cmd.Flags().IntVar(&maxTokens, "max-tokens", 0, "max tokens")

	// bind prompt flag to command
	cmd.Flags().StringVar(&prompt, "prompt", "", "prompt")

	return cmd
}

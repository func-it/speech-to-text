package speechtotext

import (
	"context"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
)

type SpeechToText struct {
	openAIKey    string
	openAIClient *openai.Client
}

func NewSpeechToText(openAIToken string) *SpeechToText {
	// XXX TODO : Assess the auth token to be a valid one
	client := openai.NewClient(openAIToken)

	return &SpeechToText{
		openAIKey:    openAIToken,
		openAIClient: client,
	}
}

func (s *SpeechToText) GetTextFromMp3AudioMessage(ctx context.Context, rawAudio io.Reader) (string, error) {
	response, err := s.openAIClient.CreateTranscription(ctx, openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "audio.mp3", // Required in both cases, even if you use reader, but the file does not have to exist
		Reader:   rawAudio,
		// Prompt:      "",
		Temperature: 0,
		// Language:   "", // Not necessary for English
		Format: openai.AudioResponseFormatJSON,
	})

	if err != nil {
		return "", fmt.Errorf("cannot create transcription: %w", err)
	}

	return response.Text, nil
}

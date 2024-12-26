package test

import (
	"context"
	"fmt"
	"github.com/func-it/go/types"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// XXX Use the configuration file
	speechToTextAddress = fmt.Sprintf("%s:%d", "localhost", 12007)
	timeout             = 10 * time.Second
)

func TestService(t *testing.T) {
	return
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial(speechToTextAddress, opts)
	if err != nil {
		t.Error(err)
	}

	client := types.NewSpeechToTextClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	rawData, err := os.ReadFile("file/input.oga")
	if err != nil {
		t.Fatalf("cannot read input test file: %v", err)
	}

	response, err := client.SpeechToText(ctx, &types.SpeechToTextRequest{
		SourceExtension: "oga",
		Data:            rawData,
	})

	if err != nil {
		t.Fatalf("cannot get response from service: %v", err)
	}

	textExpected, err := os.ReadFile("file/expected.txt")
	if err != nil {
		t.Fatalf("cannot read expected output test file: %v", err)
	}

	if response.Text != string(textExpected) {
		t.Errorf("expected:\n %s\n, got:\n %s\n", string(textExpected), response.Text)
	}

}

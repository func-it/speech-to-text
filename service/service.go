package speechToText

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/func-it/go/logi"
	"github.com/func-it/go/proto"
)

func RunService(ctx context.Context, openaiToken, listenAddress string) error {
	speechToTextService, err := NewSpeechToTextService(openaiToken)
	if err != nil {
		return logi.ErrorNReturn(err)
	}

	logi.Infof("speech-to-text is initialized")

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterSpeechToTextServer(server, speechToTextService)
	go func() {
		logi.Infof("listening on %s", listenAddress)
		logi.ServerIsReady("speech-to-text")
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()

	return nil
}

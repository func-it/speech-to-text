package audioConverter

import (
	"context"
	"log"
	"net"

	proto "github.com/func-it/speechToText/gen/service"
	"github.com/func-it/speechToText/pkg/logi"
	"google.golang.org/grpc"
)

func RunService(ctx context.Context, openaiToken, listenAddress string) error {
	speechToTextService, err := NewSpeechToTextService(openaiToken)
	if err != nil {
		return logi.ErrorNReturn(err)
	}

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterSpeechToTextServer(server, speechToTextService)
	go func() {
		logi.Infof("listening on %s", listenAddress)
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()

	return nil
}

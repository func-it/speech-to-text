package service

import (
	"context"

	"github.com/func-it/speechToText/config"
	speechToText "github.com/func-it/speechToText/service"
	"github.com/spf13/cobra"
)

func NewServiceCmd(ctx context.Context, conf *config.Conf) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service",
		Short: "run service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().PersistentPreRunE(cmd.Parent(), args)
		},
	}

	cmd.AddCommand([]*cobra.Command{
		{
			Use:   "speechToText",
			Short: "run speechToText service",
			RunE: func(cmd *cobra.Command, args []string) error {
				return speechToText.RunService(ctx, conf.OpenAi.ApiKey, conf.Services.SpeechToText.GRPC.ListenerAddr())
			},
		},
	}...)

	return cmd
}

package cmd

import (
	"context"
	"os"

	"github.com/func-it/speechToText/cmd/service"
	"github.com/func-it/speechToText/cmd/tool"
	"github.com/func-it/speechToText/config"
	"github.com/func-it/speechToText/pkg/logi"
	"github.com/spf13/cobra"
)

const version = "0.0.1"

func Exec() {
	ctx := context.Background()
	rootCmd := NewRootCmd(ctx)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

func NewRootCmd(ctx context.Context) *cobra.Command {
	conf := config.New()

	cmd := &cobra.Command{
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := conf.Load()
			if err != nil {
				panic(err)
			}

			logi.SetZap(conf.Verbose)

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&conf.Config, "config", "f", "default", "config file")
	_ = config.BindPFlag("config", cmd.PersistentFlags().Lookup("config"))

	cmd.PersistentFlags().VarP(&conf.Verbose, "verbose", "v", "set cli verbosity")
	_ = config.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))

	cmd.PersistentFlags().StringVarP(&conf.OpenAi.ApiKey, "gpt-key", "", "", "chat gpt api key")
	_ = config.BindPFlag("OpenAi.ApiKey", cmd.PersistentFlags().Lookup("gpt-key"))

	cmd.AddCommand(
		service.NewServiceCmd(ctx, conf),
		tool.NewToolCmd(ctx, conf),
	)

	return cmd
}

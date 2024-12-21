package main

import (
	"context"
	"github.com/func-it/go/logi"
	"log"

	"github.com/func-it/speech-to-text/config"
	speechtotext "github.com/func-it/speech-to-text/service"
)

func main() {
	conf, err := config.NewFromEnv()
	if err != nil {
		panic(err)
	}

	initLog(logi.InfoLevel, conf.Environment)

	err = speechtotext.RunService(context.Background(),
		conf.OpenAIApiKey,
		conf.GetAddress(),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func initLog(level logi.Level, mode string) {
	switch mode {
	case "development", "dev":
		logi.SetZap(level, logi.DevelopmentMode)
	default:
		logi.SetZap(level, logi.ProductionMode)
	}
}

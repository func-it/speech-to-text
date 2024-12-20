package main

import (
	"context"
	"log"

	"github.com/func-it/speech-to-text/config"
	speechtotext "github.com/func-it/speech-to-text/service"
)

func main() {
	conf, err := config.NewFromEnv()
	if err != nil {
		panic(err)
	}

	err = speechtotext.RunService(context.Background(),
		conf.OpenAIApiKey,
		conf.GetAddress(),
	)
	if err != nil {
		log.Fatal(err)
	}
}

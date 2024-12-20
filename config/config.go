package config

import (
	"fmt"

	"github.com/Netflix/go-env"
)

type Conf struct {
	OpenAIApiKey string `env:"LLMS_OPENAI_APIKEY,required=true"`
	GrpcPort     int    `env:"SERVICES_SPEECHTOTEXT_GRPC_ADDRESS,default=12007"`
}

func (c Conf) GetAddress() string {
	return fmt.Sprintf(":%d", c.GrpcPort)
}

func NewFromEnv() (Conf, error) {
	var conf Conf
	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		return Conf{}, err
	}

	return conf, nil
}

package config

import (
	"fmt"

	"github.com/func-it/speechToText/pkg/logi"
)

type Conf struct {
	Services *Services        `yaml:"services"`
	OpenAi   *OpenAi          `yaml:"openai"`
	Verbose  logi.EnumVerbose `yaml:"verbose"`
	Config   string           `yaml:"config"`
}

func New() *Conf {
	return &Conf{
		OpenAi: &OpenAi{},
		Services: &Services{
			SpeechToText: NewService(),
		},
	}
}

type OpenAi struct {
	ApiKey string
}

type Services struct {
	SpeechToText *Service `yaml:"speechToText"`
}

func NewService() *Service {
	return &Service{
		GRPC: &GRPC{},
	}
}

type Service struct {
	GRPC *GRPC `yaml:"grpc"`
}

type GRPC struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (g *GRPC) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}

func (g *GRPC) ListenerAddr() string {
	return fmt.Sprintf(":%d", g.Port)
}

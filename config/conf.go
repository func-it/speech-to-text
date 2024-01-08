package config

import (
	"fmt"

	"github.com/func-it/speechToText/pkg/logi"
)

type Conf struct {
	Services *Services
	OpenAi   *OpenAi
	Verbose  logi.EnumVerbose
	Config   string
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
	SpeechToText *Service
}

func NewService() *Service {
	return &Service{
		GRPC: &GRPC{},
	}
}

type Service struct {
	GRPC *GRPC
}

type GRPC struct {
	Host string
	Port int
}

func (g *GRPC) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}

func (g *GRPC) ListenerAddr() string {
	return fmt.Sprintf(":%d", g.Port)
}

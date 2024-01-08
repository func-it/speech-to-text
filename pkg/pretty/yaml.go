package pretty

import (
	_ "embed"
	"encoding/json"

	"github.com/ghodss/yaml"
	"github.com/zyedidia/highlight"
)

var (
	//go:embed syntax_yaml.yaml
	syntaxYamlYaml string
)

func PrintYaml(input any) {
	defYaml, err := highlight.ParseDef([]byte(syntaxYamlYaml))
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}

	b, err = yaml.JSONToYAML(b)
	if err != nil {
		panic(err)
	}

	print(string(b), defYaml)
}

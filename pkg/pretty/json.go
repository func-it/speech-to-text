package pretty

import (
	_ "embed"
	"fmt"

	"github.com/zyedidia/highlight"
)

var (
	//go:embed syntax_json.yaml
	syntaxJsonYaml string
)

func PrintJson(input string) {
	// Parse it into a `*highlight.Def`
	defJson, err := highlight.ParseDef([]byte(syntaxJsonYaml))
	if err != nil {
		fmt.Println(err)
		return
	}

	print(input, defJson)
}

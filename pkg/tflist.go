package tfstateviz

import (
	"github.com/kishaningithub/tf-state-viz/pkg/parser"
	"io"
)

func ShowView(stateJsonReader io.Reader, viewType string) (string, error) {
	resources, err := parser.NewTerraformStateJsonParser(stateJsonReader).Parse()
	if err != nil {
		return "", err
	}
	if viewType == ViewTypeBurnDownList {
		return BurnDown(resources), nil
	}
	if viewType == ViewTypeBurnDownListDetailed {
		return BurnDownDetailed(resources), nil
	}
	return "", nil
}

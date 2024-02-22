package tfstateviz

import (
	"github.com/kishaningithub/tf-state-viz/pkg/internal/parser"
	"io"
)

const (
	ViewTypeBurnDownList = "burndownlist"
)

func ShowView(stateJsonReader io.Reader, viewType string) (string, error) {
	resources, err := parser.NewTerraformStateJsonParser(stateJsonReader).Parse()
	if err != nil {
		return "", err
	}
	if viewType == ViewTypeBurnDownList {
		return BurnDownDetailed(resources), nil
	}
	return "", nil
}

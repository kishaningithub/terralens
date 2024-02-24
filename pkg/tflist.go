package tfstateviz

import (
	"github.com/kishaningithub/terralens/pkg/parser"
	"github.com/kishaningithub/terralens/pkg/view"
	"io"
)

func ShowView(stateJsonReader io.Reader, viewType string) (string, error) {
	resources, err := parser.NewTerraformStateJsonParser(stateJsonReader).Parse()
	if err != nil {
		return "", err
	}
	renderedView, err := view.Render(viewType, resources)
	if err != nil {
		return "", err
	}
	return renderedView, nil
}

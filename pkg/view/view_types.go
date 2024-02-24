package view

import (
	"fmt"
	"github.com/kishaningithub/terralens/pkg/parser"
	"github.com/kishaningithub/terralens/pkg/view/burndown"
	"github.com/kishaningithub/terralens/pkg/view/burndowndetailed"
)

const (
	BurnDownListDetailed = "burn_down_list_detailed"
	BurnDownList         = "burn_down_list"
)

func Render(viewType string, resources []parser.TerraformResource) (string, error) {
	switch viewType {
	case BurnDownList:
		return burndown.BurnDown(resources), nil
	case BurnDownListDetailed:
		return burndowndetailed.BurnDownDetailedView(resources), nil
	}
	return "", fmt.Errorf("invalid display type %q", viewType)
}

func SupportedViewTypes() []string {
	return []string{BurnDownList, BurnDownListDetailed}
}

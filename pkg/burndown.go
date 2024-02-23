package tfstateviz

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kishaningithub/terralens/pkg/parser"
	"strings"
)

func BurnDown(resources []parser.TerraformResource) string {
	groupedBurnDownList := make(map[string]int)
	for _, resource := range resources {
		moduleName := findResourceIdentifier(resource.Address)
		_, exists := groupedBurnDownList[moduleName]
		if exists {
			groupedBurnDownList[moduleName] = groupedBurnDownList[moduleName] + 1
			continue
		}
		groupedBurnDownList[moduleName] = 1
	}

	tableWriter := table.NewWriter()
	tableWriter.SortBy([]table.SortBy{
		{
			Number: 2,
			Mode:   table.DscNumeric,
		},
		{
			Number: 1,
			Mode:   table.Asc,
		},
	})
	tableWriter.AppendHeader(table.Row{"Resource name", "No. of resources"})
	tableWriter.SetAutoIndex(true)
	for key, value := range groupedBurnDownList {
		tableWriter.AppendRow(table.Row{key, value})
	}
	return tableWriter.RenderMarkdown()
}

func findResourceIdentifier(address string) string {
	if strings.HasPrefix(address, "module.") {
		return strings.Join(strings.Split(address, ".")[0:2], ".")
	}
	return strings.Split(address, ".")[0]
}

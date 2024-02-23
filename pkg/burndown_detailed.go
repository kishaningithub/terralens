package tfstateviz

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kishaningithub/terralens/pkg/parser"
	"strings"
)

type BurnDownOverviewData struct {
	ModuleName   string
	ResourceType string
}

func BurnDownDetailed(resources []parser.TerraformResource) string {
	var burnDownLst []BurnDownOverviewData
	for _, resource := range resources {
		burnDownLst = append(burnDownLst, BurnDownOverviewData{
			ModuleName:   findModuleName(resource.Address),
			ResourceType: resource.Type,
		})
	}
	groupedBurnDownList := make(map[BurnDownOverviewData]int)
	for _, burnDownData := range burnDownLst {
		_, exists := groupedBurnDownList[burnDownData]
		if exists {
			groupedBurnDownList[burnDownData] = groupedBurnDownList[burnDownData] + 1
			continue
		}
		groupedBurnDownList[burnDownData] = 1
	}
	tableWriter := table.NewWriter()
	tableWriter.SortBy([]table.SortBy{
		{
			Name: "Module",
			Mode: table.Asc,
		},
		{
			Name: "Count",
			Mode: table.DscNumeric,
		},
	})
	tableWriter.AppendHeader(table.Row{"Module", "Resource type", "Count"})
	for key, value := range groupedBurnDownList {
		tableWriter.AppendRow(table.Row{key.ModuleName, key.ResourceType, value})
	}
	return tableWriter.RenderMarkdown()
}

func findModuleName(address string) string {
	if strings.HasPrefix(address, "module.") {
		return strings.Split(address, ".")[1]
	}
	return "root"
}

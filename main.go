package main

import (
	"fmt"
	tfstateviz "github.com/kishaningithub/terralens/pkg"
	"github.com/kishaningithub/terralens/pkg/view"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var Version = "dev"

func main() {
	var viewType string
	var rootCmd = &cobra.Command{
		Use:     "terralens [flags]",
		Short:   "See your Terraform state world with clarity and precision - TerraLens, your visual command center for infrastructure insight!",
		Version: Version,
		Example: `
## Show burn down list
terraform show -json | terralens --type burn_down_list_detailed

## Show default view (currently burn_down_list)
terraform show -json | terralens
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			renderedView, err := tfstateviz.ShowView(os.Stdin, viewType)
			if err != nil {
				return err
			}
			fmt.Println(renderedView)
			return nil
		},
	}
	usage := fmt.Sprintf("must be one of %s", strings.Join(view.SupportedViewTypes(), ","))
	rootCmd.Flags().StringVarP(&viewType, "type", "t", view.BurnDownList, usage)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

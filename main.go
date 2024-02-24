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
	var displayType string
	var rootCmd = &cobra.Command{
		Use:     "terralens [flags]",
		Short:   "See your Terraform state world with clarity and precision - TerraLens, your visual command center for infrastructure insight!",
		Version: Version,
		Example: `
## Default usage (shows burn_down_list)
terraform show -json | terralens

## Show detailed burn down list
terraform show -json | terralens --display burn_down_list_detailed
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			renderedView, err := tfstateviz.ShowView(os.Stdin, displayType)
			if err != nil {
				return err
			}
			fmt.Println(renderedView)
			return nil
		},
	}
	usage := fmt.Sprintf("must be one of %s", strings.Join(view.SupportedViewTypes(), ","))
	rootCmd.Flags().StringVarP(&displayType, "display", "d", view.BurnDownList, usage)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

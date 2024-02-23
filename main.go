package main

import (
	"fmt"
	tfstateviz "github.com/kishaningithub/terralens/pkg"
	"github.com/spf13/cobra"
	"os"
)

var Version = "dev"

func main() {
	var viewType string
	var rootCmd = &cobra.Command{
		Use:     "terralens [flags]",
		Short:   "Visualize your terraform state in all sorts of ways",
		Version: Version,
		Example: `
## Show burn down list
terraform show -json | terralens --type burndownlist

## Show default view (currently burndownlist)
terraform show -json | terralens
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			view, err := tfstateviz.ShowView(os.Stdin, viewType)
			if err != nil {
				return err
			}
			fmt.Println(view)
			return nil
		},
	}
	usage := fmt.Sprintf("must be one of %q %q", tfstateviz.ViewTypeBurnDownList, tfstateviz.ViewTypeBurnDownListDetailed)
	rootCmd.Flags().StringVarP(&viewType, "type", "t", tfstateviz.ViewTypeBurnDownList, usage)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

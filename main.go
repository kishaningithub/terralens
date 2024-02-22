package main

import (
	"fmt"
	tfstateviz "github.com/kishaningithub/tf-state-viz/pkg"
	"github.com/spf13/cobra"
	"os"
)

var Version = "dev"

func main() {
	var viewType string
	var rootCmd = &cobra.Command{
		Use:     "tf-state-viz [flags] address",
		Short:   "Visualize your terraform state in all sorts of ways",
		Version: Version,
		Example: `
## Show burn down list
terraform show -json | tf-state-viz --type burndownlist

## Show default view (currently burndownlist)
terraform show -json | tf-state-viz
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
	rootCmd.Flags().StringVarP(&viewType, "type", "t", tfstateviz.ViewTypeBurnDownList, fmt.Sprintf("must be one of %q", tfstateviz.ViewTypeBurnDownList))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

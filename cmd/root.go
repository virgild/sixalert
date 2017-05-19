package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Long:  "sixalert - Your TTC alerts",
	Short: "sixalert",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}

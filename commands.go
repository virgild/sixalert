package main

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Long:  "sixalert - Your TTC alerts",
	Short: "sixalert",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var printCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch and print current alerts",
	Run: func(cmd *cobra.Command, args []string) {
		f := NewFetcher()
		if err := f.FetchCurrent(); err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching: %s\n", err)
			os.Exit(1)
		}
		for n, item := range f.CurrentAlerts() {
			fmt.Printf("%d. %s\n", n+1, item)
		}
	},
}

func init() {
	RootCmd.AddCommand(printCmd)
}

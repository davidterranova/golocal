package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "golocal",
}

func Execute() error {
	return rootCmd.Execute()
}

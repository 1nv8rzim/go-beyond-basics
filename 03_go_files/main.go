package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "app is a CLI application",
}

func main() {
	rootCmd.Execute()
}

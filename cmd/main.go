package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "git-harvest",
		Short: "Clone all repositories from a GitHub user or organization",
		Long: `A CLI tool to clone all GitHub repositories from a specified user or organization.`,
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

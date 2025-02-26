package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

func main() {
	var(
		token string
		target string
		isOrg bool
		destDir string
	)

	var cloneCmd = &cobra.Command{
		Use: "clone",
		Short: "Clone all repositories from a GitHub user or organization",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cloning repositories for %s (Organization: %v)\n", target, isOrg)
		},
	}

	cloneCmd.Flags().StringVarP(&token, "token", "t", "", "GitHub API token (required)")
	cloneCmd.Flags().StringVarP(&target, "target", "T", "", "GitHub user or organization name (required)")
	cloneCmd.Flags().BoolVarP(&isOrg, "org", "o", false, "Target is an organization")
	cloneCmd.Flags().StringVarP(&destDir, "dest", "d", ".", "Destination directory for cloning")

	cloneCmd.MarkFlagRequired("token")
	cloneCmd.MarkFlagRequired("target")

	var rootCmd = &cobra.Command{
		Use: "git-harvest",
		Short: "Clone all repositories from a GitHub user or organization",
		Long: `A CLI tool to clone all GitHub repositories from a specified user or organization.`,
	}

	rootCmd.AddCommand(cloneCmd)
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

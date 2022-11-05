package cmd

import (
	"os"
	"solution-generator/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "solution-generator",
	Short: "Cli tool for save solution templates",
	Long:  `Cli tool for save solution templates, and generate new project from template`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	if !internal.IsGeneratorDirExist() {
		internal.MakeInstall()
	}
}

package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"solution-generator/internal"
)

var (
	solutionName string
	solutionPath string
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save your solution to storage",
	Long:  `Save solution to generator directory`,
	Run: func(cmd *cobra.Command, args []string) {
		if solutionPath == "" {
			solutionPath = internal.GetCurrentDirPath()
		}

		internal.SaveSolution(solutionPath, solutionName)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.Flags().StringVarP(&solutionName, "name", "n", "", "New solution name")
	if err := saveCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("%e", err)
	}

	saveCmd.Flags().StringVarP(&solutionPath, "path", "p", "", "Path to solution")
}

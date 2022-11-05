package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
	"solution-generator/internal"
	"strings"
)

var (
	savedSolutionName   string
	pathToApplySolution string
	applyDirName        string
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply project template by name",
	Long:  `Apply project template by solution name`,
	Run: func(cmd *cobra.Command, args []string) {
		if !internal.SolutionExist(savedSolutionName) {
			log.Fatalf("solution '%s' not exist", savedSolutionName)
		}

		if pathToApplySolution == "" {
			pathToApplySolution = internal.GetCurrentDirPath()
		}

		if applyDirName == "" {
			pathParts := strings.Split(internal.GetSolutionSavePath(savedSolutionName), "/")
			applyDirName = pathParts[len(pathParts)-1]
		}

		pathToApplySolution := filepath.Join(pathToApplySolution, applyDirName)

		internal.ApplySolution(internal.GetSolutionSavePath(savedSolutionName), pathToApplySolution)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringVarP(&savedSolutionName, "name", "n", "", "Solution name to apply")
	if err := applyCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("name is required argument")
	}

	applyCmd.Flags().StringVarP(&pathToApplySolution, "path", "p", "", "Path to apply")

	applyCmd.Flags().StringVarP(&applyDirName, "dir", "d", "", "Dir name for new project")
}

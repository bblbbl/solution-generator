package cmd

import (
	"log"
	"solution-generator/internal"

	"github.com/spf13/cobra"
)

var (
	solutionNameToDelete string
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete solution from storage",
	Long:  `Delete solution from storage`,
	Run: func(cmd *cobra.Command, args []string) {
		if !internal.SolutionExist(solutionNameToDelete) {
			log.Fatalf("solution '%s' is not exit", solutionNameToDelete)
		}

		internal.DeleteSolution(solutionNameToDelete)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&solutionNameToDelete, "name", "n", "", "Solution name to delete")
	if err := deleteCmd.MarkFlagRequired("name"); err != nil {
		log.Fatal("name is required argument")
	}
}

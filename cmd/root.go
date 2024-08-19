package cmd

import (
	"TaskTracker/internal/storages/json_storage"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var JsonStorage *json_storage.JsonStorage

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "TaskTracker",
	Short: "CLI Application for task tracking",
	Long: `This application is a CLI application for task tracking.
This application is a tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	JsonStorage, _ = json_storage.New("storage.json")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

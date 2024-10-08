package cmd

import (
	"TaskTracker/internal/storages/json_storage"
	"os"

	"github.com/spf13/cobra"
)

var JsonStorage *json_storage.JsonStorage

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "TaskTracker",
	Short: "CLI Application for task tracking",
	Long: `This application is a CLI application for task tracking.
This application is a tool for managing your tasks.
You can add, delete, update and view your tasks using special commands.`,
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
}

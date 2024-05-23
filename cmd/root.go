package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "file-organizer",
	Short: "A cli tool organize files",
	Long: `A Command Line Interface (CLI) tool to organize files in a directory based on file type, date, or custom rules.`,
} 

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

func init() {
    
}
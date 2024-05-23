package cmd

import (
	"fmt"

	"github.com/Ricky004/forga/organizer"

	"github.com/spf13/cobra"
)

var dir string
var configPath string
var date bool
var prefix string
var suffix string
var startNumber int

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize files in a directory by type",
	Long: `Organize files in the specified directory
	by their type, such as images, documents, videos, and music.`,
	Run: func(cmd *cobra.Command, args []string) {
		if dir == "" {
			fmt.Println("Please specify a directory using the --dir flag.")
			return
		}
		if date {
			err := organizer.OrganizeByDate(dir)
			if err != nil {
				fmt.Printf("Error organizing files: %v\n", err)
			}
		} else {
			err := organizer.CategorizeByType(dir, configPath)
			if err != nil {
				fmt.Printf("Error organizing files: %v\n", err)
			}
		}
		if prefix != "" || suffix != "" || startNumber != 0 {
			err := organizer.BulkRenaming(dir, prefix, suffix, startNumber)
			if err != nil {
				fmt.Printf("Error organizing files: %v\n", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
	organizeCmd.Flags().StringVar(&dir, "dir", "", "Directory to organize")
	organizeCmd.Flags().StringVar(&configPath, "config", "", "Path to the configuration file (optional)")
	organizeCmd.Flags().BoolVar(&date, "date", false, "Organize by date")
	organizeCmd.Flags().StringVar(&prefix, "prefix", "", "Prefix to add to file names")
	organizeCmd.Flags().StringVar(&suffix, "suffix", "", "Suffix to add to file names")
	organizeCmd.Flags().IntVar(&startNumber, "start-number", 0, "Starting number for sequential renaming")
}

package cmd

import (
	"fmt"

	"github.com/Ricky004/forga/organizer"

	"github.com/spf13/cobra"
)

var (
	dir             string
	configPath      string
	date            bool
	prefix          string
	suffix          string
	startNumber     int
	remove          bool
	relocate        string
)

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
		if remove || relocate != "" {
			duplicates, err := organizer.DetectDuplicates(dir)
			if err != nil {
				fmt.Printf("Error detecting duplicates: %v\n", err)
				return
			}
			if remove {
				err = organizer.RemoveDuplicates(duplicates)
				if err != nil {
					fmt.Printf("Error removing duplicates: %v\n", err)
					return
				}
			} else if relocate != "" {
				err = organizer.RelocateDuplicates(duplicates, relocate)
				if err != nil {
					fmt.Printf("Error detecting duplicates: %v\n", err)
					return
				}
			}
		}
		
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
	organizeCmd.Flags().StringVar(&dir, "dir", "", "Directory to organize")
	organizeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to the configuration file (optional)")
	organizeCmd.Flags().BoolVarP(&date, "date", "d", false, "Organize by date")
	organizeCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "Prefix to add to file names")
	organizeCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "Suffix to add to file names")
	organizeCmd.Flags().IntVarP(&startNumber, "start-number", "n", 0, "Starting number for sequential renaming")
	organizeCmd.Flags().BoolVar(&remove, "remove", false, "Remove duplicate files")
	organizeCmd.Flags().StringVar(&relocate, "relocate", "", "Directory to relocate duplicate files")
}

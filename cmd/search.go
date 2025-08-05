package cmd

import (
	"fmt"

	"github.com/Ricky004/forg/organizer"

	"github.com/spf13/cobra"
)

var (
    
	searchName      string
	searchExtension string
	minSize         string
	maxSize         string
	beforeDate      string
	afterDate       string
)

var searchCmd = &cobra.Command{
    Use:   "search",
    Short: "Search and filter files",
    Long:  `Search and filter files based on name, extension, size, and modification date.`,
	Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        err := organizer.SearchFiles(args, searchName, searchExtension, beforeDate, afterDate, minSize, maxSize)
        if err != nil {
            fmt.Println("Error:", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(searchCmd)
    searchCmd.Flags().StringVarP(&searchName, "name", "n", "", "Filter by file name")
    searchCmd.Flags().StringVarP(&searchExtension, "extension", "e", "", "Filter by file extension")
    searchCmd.Flags().StringVar(&minSize, "min-size", "", "Minimum file size (e.g., 3b, 50kb, 100mb)")
    searchCmd.Flags().StringVar(&maxSize, "max-size", "", "Minimum file size (e.g., 3b, 50kb, 100mb)")
    searchCmd.Flags().StringVar(&beforeDate, "before", "", "Filter files modified before this date (YYYY-MM-DD)")
    searchCmd.Flags().StringVar(&afterDate, "after", "", "Filter files modified after this date (YYYY-MM-DD)")
}
package cmd

import (
	"errors"
	"fmt"
	"slices"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Removes non-existent entries in the clipboard",
	Long: `clean removes all non-existent entries in the clipboard.
non-existent entries in this case are entries that have links to files that
may have been deleted or moved to another location.
clean helps to make space for newer entries in the clipboard.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 0 {
			return errors.New("clean doesn't require any arguments")
		}

		err := cleanClipboard()
		if err != nil {
			return errors.New("failed to clean the clipboard")
		}

		return nil
	},
}

func cleanClipboard() error {
	var existingEntries []string
	entries := getLinesFrom(0)
	if entries == nil {
		return fmt.Errorf("there are no items in the clipboard. add an item with 'yyt add'…")
	}
	missingEntries, _ := sortMissingEntries(entries)

	if missingEntries == nil {
		printSuccess("the clipboard is clean; no dead files.")
		return nil
	}

	// add all live links to the slice that will be written
	for _, e := range entries {
		if !slices.Contains(missingEntries, e) {
			existingEntries = append(existingEntries, e.filePath)
		}
	}

	err := writeToFile(existingEntries)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	printSuccess("the following files have been cleaned successfully")
	printFilesName(missingEntries)

	return nil
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

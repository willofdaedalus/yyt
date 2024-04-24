/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cleanClipboard()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("yyt: clean doesn't require any arguments")
		}
	},
}

func cleanClipboard() error {
	var liveLinks []string
	entries := getLinesFrom(0)
	if entries == nil {
		return fmt.Errorf(
			"yyt: there are no items in the clipboard. add an item with 'yyt add'...")
	}
	missingEntries, _ := sortMissingEntries(entries)

	if missingEntries == nil {
		fmt.Println("yyt: clipboard is clean; no missing files.")
		return nil
	}

	// add all live links to the slice that will be written
	for _, e := range entries {
		if !slices.Contains(missingEntries, e) {
			liveLinks = append(liveLinks, e.filePath)
		}
	}

	temp, _ := os.CreateTemp("", "yyt-*")
	defer os.Remove(temp.Name())

	// write all valid entries containing updated files to the tempfile
	// validFiles filters out the invalid files, prints them and returns
	// all valid files
	if _, err := temp.Write([]byte(strings.Join(
		liveLinks, "\n") + "\n")); err != nil {
		return fmt.Errorf("error writing to temp file: %w", err)
	}

	defer func() {
		fmt.Println(
			"yyt: the following files have been cleaned successfully")

		for _, f := range missingEntries {
			fmt.Println(f.fileName)
		}
	}()

	// rename and replace the old clipboard file
	os.Rename(temp.Name(), clipboardLocation)
	return nil
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"slices"

	// "github.com/google/uuid"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add files to the clipboard",
	Long:  `add files to the clipoard that allows you to keep track of them`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No files to add")
			return
		}

		fakes, skippedEntries, uniqueFiles := filterDuplicates(args)

		if uniqueFiles != nil {
			err := addFile(uniqueFiles)
			if err != nil {
				fmt.Print(err)
			}
		}

		if skippedEntries != nil {
			fmt.Println(
				"\nyyt: the following args are already in the clipboard; skipped.")

			for i, value := range skippedEntries {
				fmt.Printf("%v. %s\n", i, value)
			}
		}

		if fakes != nil {
			fmt.Println(
				"\nyyt: the following args are either not files or are directories; skipped.")

			for i, value := range fakes {
				fmt.Printf("%v. %s\n", i, value)
			}
		}
	},
}

// addFile adds files to the clipboard
func addFile(files []ClipboardEntry) error {
	var allFilePaths []string
    message := "yyt: the following files have been added to successfully"
	currentSize, err := fileSize()

	if err != nil {
		return fmt.Errorf("error getting file size: %w", err)
	}

	// clipboard is full
	if currentSize >= maxFiles {
		// calculate the number of lines to keep
		var lastLine = 0
		fileLen := len(files)
		if maxFiles > fileLen {
			lastLine = fileLen
		} else {
			lastLine = fileLen - maxFiles
		}

		// get the last lines from the clipboard that would allow enough
		// space for the new entries into the clipboard and append the new
		// entries to the slice before finally writing it the clipboard
		var oldLines []ClipboardEntry = getLinesFrom(lastLine)
		for i := lastLine - 1; i < len(files); i++ {
			oldLines = append(oldLines, files[i])
		}

		// create a temp file to write to
		for _, line := range oldLines {
			allFilePaths = append(allFilePaths, line.filePath)
		}

        writeToFile(message, allFilePaths, files)

		return nil
	}

	f, err := os.OpenFile(clipboardLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v\n", err)
	}
	defer f.Close()

	for _, line := range files {
		allFilePaths = append(allFilePaths, line.filePath)
	}

	for _, fileEntry := range allFilePaths {
		if _, err := f.Write([]byte(fileEntry + "\n")); err != nil {
			return fmt.Errorf("error getting file size: %w", err)
		}
	}

    defer func() {
        fmt.Println(message)

        for _, f := range files {
            fmt.Println(f.fileName)
        }
    }()

	return nil
}

// removes duplicates from being added to the clipboard
// returns a slice of fake files, a slice of skipped files and a slice of unique files
func filterDuplicates(userArgs []string) ([]string, []string, []ClipboardEntry) {
	var (
		nonDuplicates  []ClipboardEntry
		skippedEntries []string
	)

	fakes, uniqueFiles := makeEntriesSlice(userArgs)
	existingEntries := getLinesFrom(0)

	// check if there are already entries in the clipboard
	if existingEntries != nil {
		for _, file := range uniqueFiles {
			if !slices.Contains(existingEntries, file) {
				nonDuplicates = append(nonDuplicates, file)
			} else {
				skippedEntries = append(skippedEntries, file.fileName)
			}
		}
	} else {
        // if the file doesn't exist
		for _, entry := range uniqueFiles {
			if !slices.Contains(nonDuplicates, entry) {
				nonDuplicates = append(nonDuplicates, entry)
			}
		}
	}

	return fakes, skippedEntries, nonDuplicates
}

func init() {
	rootCmd.AddCommand(addCmd)
}

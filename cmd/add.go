/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	// "github.com/google/uuid"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add files to the clipboard",
	Long:  `add files to the clipboard that allows you to keep track of them`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("no files to add")
		}

		fakes, skippedEntries, uniqueFiles := filterDuplicates(args)

		if uniqueFiles != nil {
			err := addFile(uniqueFiles)
			if err != nil {
				cmd.SilenceUsage = true // no need to display Usage when a real error occurs
				return fmt.Errorf("adding files: %w", err)
			}
		}

		if skippedEntries != nil {
			fmt.Println(
				"yyt: the following args are already in the clipboard; skipped.")

			for i, value := range skippedEntries {
				fmt.Printf("%v. %s\n", i, value)
			}
		}

		if fakes != nil {
			fmt.Println(
				"yyt: the following args are either not files or are directories; skipped.")

			for i, value := range fakes {
				fmt.Printf("%v. %s\n", i, value)
			}
		}

		return nil
	},
}

// addFile adds files to the clipboard
func addFile(files []ClipboardEntry) error {
	var allFilePaths []string
	currentSize, err := fileSize()
	if err != nil {
		return fmt.Errorf("error getting file size: %w", err)
	}
	message := "the following files have been added to successfully"

	// clipboard is full
	if currentSize >= maxFiles {
		// calculate the number of lines to keep
		lastLine := 0
		fileLen := len(files)
		if maxFiles > fileLen {
			lastLine = fileLen
		} else {
			lastLine = fileLen - maxFiles
		}

		// get the last lines from the clipboard that would allow enough
		// space for the new entries into the clipboard and append the new
		// entries to the slice before finally writing it the clipboard
		oldLines := getLinesFrom(lastLine)
		for i := lastLine - 1; i < len(files); i++ {
			oldLines = append(oldLines, files[i])
		}

		// create a temp file to write to
		for _, line := range oldLines {
			allFilePaths = append(allFilePaths, line.filePath)
		}

		err = writeToFile(allFilePaths)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}

		printSuccess(message)
		printFilesName(files)

		return nil
	}

	f, err := os.OpenFile(clipboardLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	for _, line := range files {
		allFilePaths = append(allFilePaths, line.filePath)
	}

	for _, fileEntry := range allFilePaths {
		if _, err := f.WriteString(fileEntry + "\n"); err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	printSuccess(message)
	printFilesName(files)

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

	// Check if there are already entries in the clipboard
	// there must be a better way to do this
	// put a pin in it
	if existingEntries != nil {
		for _, file := range uniqueFiles {
			duplicate := false

			for _, existingFile := range existingEntries {
				if file.filePath == existingFile.filePath {
					skippedEntries = append(skippedEntries, file.fileName)
					duplicate = true
					break
				}
			}
			if !duplicate {
				nonDuplicates = append(nonDuplicates, file)
			}
		}
	} else {
		// If the clipboard is empty, all incoming files are non-duplicates
		nonDuplicates = uniqueFiles
	}

	return fakes, skippedEntries, nonDuplicates
}

func init() {
	rootCmd.AddCommand(addCmd)
}

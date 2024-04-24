/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

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

		fakes, actualFiles := makeEntriesSlice(args)

		if actualFiles != nil {
			err := addFile(actualFiles)
			if err != nil {
				fmt.Print(err)
			}
		}

		if fakes != nil {
			fmt.Println(
				"\nyyt: the following args are either not files or are directories. skipping...")

			for i, value := range fakes {
				fmt.Printf("%v. %s\n", i, value)
			}
		}
	},
}

// addFile adds files to the clipboard
func addFile(files []ClipboardEntry) error {
	var allFilePaths []string
	currentSize, err := fileSize()
	if err != nil {
		return fmt.Errorf("error getting file size: %w", err)
	}

	// clipboard is full
	if currentSize >= maxFiles {
		// calculate the number of lines to keep
		var lastLines = 0
		fileLen := len(files)
		if maxFiles > fileLen {
			lastLines = fileLen
		} else {
			lastLines = fileLen - maxFiles
		}

		// get the last lines from the clipboard that would allow enough
		// space for the new entries into the clipboard and append the new
		// entries to the slice before finally writing it the clipboard
		var oldLines []ClipboardEntry = getLinesFrom(lastLines)
		for i := lastLines - 1; i < len(files); i++ {
			oldLines = append(oldLines, files[i])
		}

		// create a temp file to write to
		temp, _ := os.CreateTemp("", "yyt-*")
		defer os.Remove(temp.Name())

		for _, line := range oldLines {
			allFilePaths = append(allFilePaths, line.filePath)
		}

		// write all valid entries containing updated files to the tempfile
		// validFiles filters out the invalid files, prints them and returns
		// all valid files
		if _, err := temp.Write([]byte(strings.Join(
			allFilePaths, "\n") + "\n")); err != nil {
			return fmt.Errorf("error writing to temp file: %w", err)
		}

		defer func(okValues []string) {
			fmt.Println(
				"yyt: the following files have been added to successfully")

			for _, f := range files {
				fmt.Println(f.fileName)
			}
		}(allFilePaths)

		// rename and replace the old clipboard file
		os.Rename(temp.Name(), clipboardLocation)
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

	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}

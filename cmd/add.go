/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
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

        err := addFile(args)
        if err != nil {
            fmt.Print(err)
        }
	},
}

// addFile adds files to the clipboard
func addFile(files []string) error {
    // check if there are invalid files being added
    allValidFiles, err := validFiles(files)
    if err != nil {
        return err
    }

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

		var oldLines []string = getLastLines(lastLines)
		for i := lastLines - 1; i < len(files); i++ {
			oldLines = append(oldLines, files[i])
		}

		// create a temp file to write to
		temp, _ := os.CreateTemp("", "yyt-*")
		defer os.Remove(temp.Name())

        valids, _ := validFiles(oldLines)

		// write all valid entries containing updated files to the tempfile
        // validFiles filters out the invalid files, prints them and returns
        // all valid files
		if _, err := temp.Write([]byte(strings.Join(
			valids, "\n") + "\n")); err != nil {
			return fmt.Errorf("error writing to temp file: %w", err)
		}

		// rename and replace the old clipboard file
		os.Rename(temp.Name(), clipboardLocation)
		return nil
	}

	f, err := os.OpenFile(clipboardLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v\n", err)
	}
	defer f.Close()

	for _, fileEntry := range allValidFiles {
		if _, err := f.Write([]byte(fileEntry + "\n")); err != nil {
			return fmt.Errorf("error getting file size: %w", err)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

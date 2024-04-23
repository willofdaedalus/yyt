/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all files currently in the clipboard",
	Long:  `ls displays all files in the current file for inspection`,
	Run: func(cmd *cobra.Command, args []string) {
        err := listFiles(args)
        if err != nil {
            fmt.Println(err)
        }
	},
}

// helper function for listing files
func listFiles(args []string) error {
	entries := getLinesFrom(0)
    if entries == nil {
        return fmt.Errorf(
            "yyt: there are no items in the clipboard. add an item with 'yyt add'...")
    }
	// valueMappedEntries := mappedEntries(entries)
    structuredEntries := structureEntries(entries)

    // checks every one of the user's input against the clipboard's entries
	if len(args) > 0 {
        foundValue := false
        var unfoundValuesSlice []string

        for _, value := range args {
            for _, structEntry := range structuredEntries {
                if structEntry.fileName == value {
                    fmt.Printf("> %s @ %s\n", value, structEntry.filePath)
                    foundValue = true
                }
            }

            // create an array to keep all unfound values for display later
            if !foundValue {
                unfoundValuesSlice = append(unfoundValuesSlice, value)
            }
            foundValue = false
        }

        // print all args passed that are not in the buffer
        if len(unfoundValuesSlice) > 0 {
            fmt.Println(
                "\nyyt: the following passed args were not found in the clipboard")
            for _, v := range unfoundValuesSlice {
                fmt.Println(v)
            }
        }
	} else {
        // print everything if the length of args is 0
        for _, value := range structuredEntries {
            fmt.Printf("> %s @ %s\n", value.fileName, value.filePath)
        }
    }
    return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}

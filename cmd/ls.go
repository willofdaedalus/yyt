/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"strings"

	//"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all files currently in the clipboard",
	Long:  `ls displays all files in the current file for inspection`,
	Run: func(cmd *cobra.Command, args []string) {
        listFiles(args)
	},
}

// helper function for listing files
func listFiles(args []string) {
	foundUserEntry := false
	entries := getLastLines(0)
	plainFiles := rawFiles(entries)

    // checks every one of the user's input against the clipboard's entries
	if len(args) > 0 {
		for _, value := range args {
			if slices.Contains(plainFiles, value) {
				// flag to trigger when the user's args is found
				if !foundUserEntry {
					foundUserEntry = true
				}

				idxOfEntry := slices.Index(plainFiles, value)
				fmt.Printf("> %s @ %s\n", value, entries[idxOfEntry])
			} else {
                fmt.Printf(
                    "yyt: the file %q has not been added yet to the clipboard\n", value)
            }
		}
	} else {
        for _, value := range entries {
            // print everything if the length of args is 0
            pathPoints := strings.Split(value, "/")
            fmt.Printf("%s @ %s\n", pathPoints[len(pathPoints)-1], value)
        }
    }
}

func rawFiles(entries []string) []string {
	var retVals []string
	for _, v := range entries {
		file := strings.Split(v, "/")
		retVals = append(retVals, file[len(file)-1])
	}

	return retVals
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

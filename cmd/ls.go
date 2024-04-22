/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
	valueMappedEntries := mappedEntries(entries)

    // checks every one of the user's input against the clipboard's entries
	if len(args) > 0 {
		for _, value := range args {
            // check if the argument passed by the user is in the clipboard
            if _, ok := valueMappedEntries[value]; ok {
                fmt.Printf("> %s @ %s\n", value, valueMappedEntries[value])
			} else {
                fmt.Printf(
                    "yyt: the file %q has not been added to the clipboard yet\n", value)
            }
		}
	} else {
        // print everything if the length of args is 0
        for _, value := range entries {
            fmt.Printf("%s\n", value)
        }
    }
    return nil
}

// mappedEntries returns a map of the entries in the clipboard
// by the filename and its path on the file system
func mappedEntries(entries []string) map[string]string {
    returnMap := make(map[string]string)

    // generate unique id for each entry
    // when we're comparing, do a dynamic string operation where we
    // remove the id and replace it with a random string

    for _, v := range entries {
        paths := strings.Split(v, "/")
        returnMap[paths[len(paths) - 1]] = v
    }

    return returnMap
}

func init() {
	rootCmd.AddCommand(listCmd)
}

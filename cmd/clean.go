/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

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
        err :=  cleanClipboard()
        if err != nil {
            fmt.Println(err)
        }
	},
}

func cleanClipboard() error {
	entries := getLinesFrom(0)
	if entries == nil {
		return fmt.Errorf(
			"yyt: there are no items in the clipboard. add an item with 'yyt add'...")
	}
	missingEntries, _ := sortMissingEntries(entries)

    if missingEntries == nil {
        fmt.Println("yyt: clipboard is clean; no missing files.")
        return nil
    } else {
        fmt.Println("yyt: dead entries found. removing...")
    }

    return nil

}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

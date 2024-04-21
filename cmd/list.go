/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
    "slices"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all files currently in the clipboard",
	Long: `ls displays all files in the current file for inspection`,
	Run: func(cmd *cobra.Command, args []string) {
        err := listFiles(args)
        if err != nil {
            fmt.Print(err)
        }
	},
}

// unfinished
func listFiles(args []string) error {
    entries := getLastLines(0)
    for _, value := range entries {
        // print everything if the length of args is 0
        pathPoints := strings.Split(value, "/")
        fmt.Printf("%s @ %s\n\n", pathPoints[len(pathPoints) - 1], value)
    }

    return nil
}

func rawFiles(entries []string) []string {
    var retVals []string
    for _, v := range entries {
        file := strings.Split(v, "/")
        retVals = append(retVals, file[len(file) - 1])
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

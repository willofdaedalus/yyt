/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

const (
	clipboardLocation = "/tmp/yyt"
	maxFiles          = 10
)

// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	Use:   "yyt",
	Short: "Keep track of files in a clipboard-like manner.",
	Long: `yyt is a clipboard-like tool that allows you to keep track of files
that you wish to copy, move or delete later. It is useful when you are working
on the commandline and you want to keep track of files that you want to work.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	rootCmd.SetErrPrefix("yyt: error")
	return rootCmd.Execute()
}

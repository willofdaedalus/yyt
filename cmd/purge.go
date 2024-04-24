package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Removes all entries in the clipboard for a fresh start",
	Long: `purge removes all the entries in the clipboard if the user wishes
to start with a fresh clipboard without manually removing every entry with
yyt rm.`,
	Run: func(cmd *cobra.Command, args []string) {
        // simply delete the file itself
        if len(args) > 0 {
            fmt.Println("yyt: purge doesn't require any arguments")
            return
        }

        err := os.Remove(clipboardLocation)
        if err != nil {
            fmt.Println("yyt: there are no entries in the clipboard")
            return
        }

        fmt.Println("yyt: all entries have been cleared from the clipboard")
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}

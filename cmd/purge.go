package cmd

import (
	"errors"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("purge doesn't require any arguments")
		}

		// simply delete the file itself
		err := os.Remove(clipboardLocation)
		if err != nil {
			cmd.SilenceUsage = true // no need to display Usage when a real error occurs
			return fmt.Errorf("failed to remove file: %w", err)
		}

		printSuccess("all entries have been cleared from the clipboard")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}

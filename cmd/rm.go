package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove specific file(s) from the clipboard",
	Long: `rm removes any passed file names from the clipboard if they are
found. If nothing is found the command simply prints a message and exits.`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil {
			fmt.Println("yyt: no files specified. exiting...")
			return
		}

		err := removeFile(args)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func removeFile(entries []string) error {
	clipboardEntries := getLinesFrom(0)
	if clipboardEntries == nil {
		fmt.Println("yyt: there are no entries in the clipboard to remove.")
		return nil
	}

	var (
		entriesToKeep []string
		removed       []ClipboardEntry
	)

	// Create a map to store lowercase versions of user arguments for efficient lookup
	userArgsMap := make(map[string]struct{})
	for _, arg := range entries {
		userArgsMap[strings.ToLower(arg)] = struct{}{}
	}

	for _, clipboardEntry := range clipboardEntries {
		entryToLower := strings.ToLower(clipboardEntry.fileName)

		foundMatch := false
		for userArg := range userArgsMap {
			if strings.Contains(entryToLower, userArg) {
				removed = append(removed, clipboardEntry)
				foundMatch = true
				break // No need to check other user arguments for this clipboard entry
			}
		}

		if !foundMatch {
			// Only append the entry to keep if it hasn't been removed
			entriesToKeep = append(entriesToKeep, clipboardEntry.filePath)
		}
	}

	// If there are no entries left in the clipboard, consider that a purge
	// and delete the file
	if len(entriesToKeep) == 0 {
		err := os.Remove(clipboardLocation)
		if err != nil {
			fmt.Println("yyt: error removing clipboard file:", err)
			return err
		}
		fmt.Println("yyt: all entries have been cleared from the clipboard")
		return nil
	}

	message := "yyt: the following files have been removed from the clipboard"
	if err := writeToFile(message, entriesToKeep, removed); err != nil {
		fmt.Println("yyt: error writing to file:", err)
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

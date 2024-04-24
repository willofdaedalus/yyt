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
    var (
        entriesToKeep []string
        removed []ClipboardEntry
    )

    clipboardEntries := getLinesFrom(0)
    if clipboardEntries == nil {
        fmt.Println("yyt: there are no entries in the clipboard to remove.")
        return nil
    }

    for _, arg := range entries {
        userArgToLower:= strings.ToLower(arg)

        for _, clipboardEntry := range clipboardEntries {
            entryToLower := strings.ToLower(clipboardEntry.fileName)

            if strings.Contains(entryToLower, userArgToLower) {
                removed = append(removed, clipboardEntry)
            } else {
                entriesToKeep = append(entriesToKeep, clipboardEntry.filePath)
            }
        }
    }

    // if there are no entries left in the clipboard, consider that a purge
    // and delete the file
    if len(entriesToKeep) == 0 {
        err := os.Remove(clipboardLocation)
        if err != nil {
            fmt.Println("yyt: there are no entries in the clipboard")
            return nil
        }

        fmt.Println("yyt: all entries have been cleared from the clipboard")
        return nil
    }

    message := "yyt: the following files have been removed from the clipboard"
    writeToFile(message, entriesToKeep, removed)

    return nil
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

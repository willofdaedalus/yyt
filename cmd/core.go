package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
    "strings"
)

type ClipboardEntry struct {
    fileName, filePath string
}


// validateAllFilePaths returns a list of valid files in their absolute paths
// gets the absolute path of the file and checks if it exists before adding
// it to the list of valid files
func validateAllFilePaths(files []string) ([]string, error) {
	var validFiles []string

	for _, f := range files {
		file, _ := filepath.Abs(f)

		if !checkFileExists(file) {
			return validFiles, fmt.Errorf(
				"yyt: file %q doesn't exist or is not a file. cancelling...\n", f)
		} else {
			validFiles = append(validFiles, file)
		}
	}
	return validFiles, nil
}

func checkFileExists(file string) bool {
	if fileInfo, err := os.Stat(file); err == nil {
		return fileInfo.Mode().IsRegular()
	}
	return false
}

func structureEntries(entries []string) []ClipboardEntry {
    returnEntries := make([]ClipboardEntry, len(entries))

    for i, v := range entries {
        paths := strings.Split(v, "/")
        returnEntries[i].fileName = paths[len(paths) - 1]
        returnEntries[i].filePath = v
    }

    return returnEntries
}

func fileSize() (int, error) {
	f, _ := os.OpenFile(clipboardLocation, os.O_RDONLY|os.O_CREATE, 0644)
	defer f.Close()

	var lines int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning file: %w", err)
	}
	return lines, nil
}

func getLinesFrom(fromLineNumber int) []ClipboardEntry {
    var retEntries []ClipboardEntry
	// open the file
	f, err := os.Open(clipboardLocation)
	if err != nil {
		return nil
	}
	defer f.Close()

	// read the contents of the file
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

    lines = lines[fromLineNumber:]
    for _, line := range lines {
        paths := strings.Split(line, "/")
        fileName := paths[len(paths) - 1]
        returnEntry  := ClipboardEntry{fileName, line}

        retEntries = append(retEntries, returnEntry)
    }

	return retEntries
}

func makeEntriesSlice(files []string) ([]string, []ClipboardEntry) {
    var (
        fakes []string
        liveFiles []ClipboardEntry
    )

	for _, f := range files {
		file, _ := filepath.Abs(f)

		if !checkFileExists(file) {
            fakes = append(fakes, f)
		} else {
            tempEntry := ClipboardEntry {fileName: f, filePath: file}
            liveFiles = append(liveFiles, tempEntry)
		}
	}

    return fakes, liveFiles
}

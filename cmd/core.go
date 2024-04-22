package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

    "github.com/google/uuid"
)

// allValidFilePaths returns a list of valid files in their absolute paths
// gets the absolute path of the file and checks if it exists before adding
// it to the list of valid files
func allValidFilePaths(files []string) ([]string, error) {
	var validFiles []string

	for _, f := range files {
		file, _ := filepath.Abs(f)

		if !checkFileExists(file) {
			return validFiles, fmt.Errorf(
				"yyt: file %q doesn't exist or is not a file. cancelling...\n", f)
		} else {
            file = appendUniqueID(file)
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

func appendUniqueID(file string) string {
    uuid := uuid.New()
    return fmt.Sprintf("%s-%s", file, uuid.String())
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

func getLinesFrom(fromLineNumber int) []string {
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

	retval := lines[fromLineNumber:]
	return retval
}

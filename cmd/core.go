package cmd

import (
	"bufio"
	"fmt"
    "path/filepath"
	"os"
)

// validFiles returns a list of valid files
func validFiles(files []string) ([]string, error) {
	var validFiles []string
	for _, f := range files {
        file, _ := filepath.Abs(f)
		if fileInfo, err := os.Stat(file); err == nil {
			// in the future if we wish to include directories, it should
			// be an easy change and should be done here
			if fileInfo.Mode().IsRegular() {
				validFiles = append(validFiles, file)
			}
		} else {
            // prompt user a file doesn't exist
            return validFiles, fmt.Errorf(
                "yyt: file %q doesn't exist or is not a file. cancelling...\n", f)
        }
	}
	return validFiles, nil
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

func getLastLines(fromLineNumber int) []string {
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

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AddFile adds files to the clipboard
func AddFile(files []string) error {
	currentSize, err := fileSize()
	if err != nil {
		return fmt.Errorf("error getting file size: %w", err)
	}
	// clipboard is full
	if currentSize >= maxFiles {
        // calculate the number of lines to keep
        var lastLines = 0
        fileLen := len(files)
        if maxFiles > fileLen {
            lastLines = fileLen
        } else {
            lastLines = fileLen - maxFiles
        }

		var newLines []string = getLastLines(lastLines)
		for i := lastLines - 1; i < len(files); i++ {
			newLines = append(newLines, files[i])
		}

        // create a temp file to write to
		temp, _ := os.CreateTemp("", "yyt-*")
		defer os.Remove(temp.Name())

		if _, err := temp.Write([]byte(strings.Join(newLines, "\n"))); err != nil {
			return fmt.Errorf("error writing to temp file: %w", err)
		}

		os.Rename(temp.Name(), clipboardLocation)
		return nil
	}

	f, err := os.OpenFile(clipboardLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v\n", err)
	}
	defer f.Close()
	for _, fileEntry := range files {
		if _, err := f.Write([]byte(fileEntry + "\n")); err != nil {
			return fmt.Errorf("error getting file size: %w", err)
		}
	}

	return nil
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

func getLastLines(fromLineNumber int) (retval []string) {
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

    retval = lines[fromLineNumber:]
	return
}

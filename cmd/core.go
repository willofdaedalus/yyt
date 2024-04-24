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

func checkFileExists(file string) bool {
	if fileInfo, err := os.Stat(file); err == nil {
		return fileInfo.Mode().IsRegular()
	}
	return false
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
		fileName := paths[len(paths)-1]
		returnEntry := ClipboardEntry{fileName, line}

		retEntries = append(retEntries, returnEntry)
	}

	return retEntries
}

func makeEntriesSlice(files []string) ([]string, []ClipboardEntry) {
	var (
		fakes     []string
		liveFiles []ClipboardEntry
	)

	for _, f := range files {
		file, _ := filepath.Abs(f)

		if !checkFileExists(file) {
			fakes = append(fakes, f)
		} else {
			tempEntry := ClipboardEntry{fileName: f, filePath: file}
			liveFiles = append(liveFiles, tempEntry)
		}
	}

	return fakes, liveFiles
}

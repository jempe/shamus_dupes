package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileEntry struct {
	Shasum string
	Path   string
}

var totalDuplicates int

// parseLine takes a line from the file and splits it into shasum and file path
func parseLine(line string) FileEntry {
	parts := strings.Fields(line)
	return FileEntry{
		Shasum: parts[0],
		Path:   strings.Join(parts[1:], " "), // handle file paths with spaces
	}
}

// compareFiles reads two sorted files and prints shasums that match, along with their paths
func compareFiles(file1, file2 string, action string) error {
	f1, err := os.Open(file1)
	if err != nil {
		return err
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return err
	}
	defer f2.Close()

	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)

	var entry1, entry2 FileEntry

	// Initial scan of both files
	hasMore1 := scanner1.Scan()
	if hasMore1 {
		entry1 = parseLine(scanner1.Text())
	}

	hasMore2 := scanner2.Scan()
	if hasMore2 {
		entry2 = parseLine(scanner2.Text())
	}

	// Iterate through both files, comparing shasums
	for hasMore1 && hasMore2 {
		if entry1.Shasum == entry2.Shasum {
			if action == "duplicate" {
				fmt.Printf("%s is a duplicate of %s\n", entry1.Path, entry2.Path)
			} else if action == "remove" {
				fmt.Printf("rm \"%s\"\n", entry1.Path)
			}

			totalDuplicates++

			// Move both scanners to the next line
			hasMore1 = scanner1.Scan()
			if hasMore1 {
				entry1 = parseLine(scanner1.Text())
			}

			hasMore2 = scanner2.Scan()
			if hasMore2 {
				entry2 = parseLine(scanner2.Text())
			}
		} else if entry1.Shasum < entry2.Shasum {
			if action == "unique" {
				fmt.Printf("%s is unique\n", entry1.Path)
			}
			// Move file 1 forward
			hasMore1 = scanner1.Scan()
			if hasMore1 {
				entry1 = parseLine(scanner1.Text())
			}
		} else {
			if action == "unique" {
				fmt.Printf("%s is unique\n", entry2.Path)
			}

			// Move file 2 forward
			hasMore2 = scanner2.Scan()
			if hasMore2 {
				entry2 = parseLine(scanner2.Text())
			}
		}
	}

	return nil
}

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <action> <file1> <file2>\n", os.Args[0])
		fmt.Println("The action can be 'duplicate', 'unique', 'remove'")
		return
	}

	action := os.Args[1]
	file1 := os.Args[2]
	file2 := os.Args[3]

	if action == "unique" {
		fmt.Println("Unique files")
	} else if action == "duplicate" {
		fmt.Println("Duplicate files")
	} else if action == "remove" {
		fmt.Println("Remove duplicates")
	} else {
		fmt.Println("Invalid action")
		return
	}

	err := compareFiles(file1, file2, action)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("#Total duplicates: %d\n", totalDuplicates)
}

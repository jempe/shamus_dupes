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
func compareFiles(fileToRemove, fileToKeep string, action string) error {
	fToRemove, err := os.Open(fileToRemove)
	if err != nil {
		return err
	}
	defer fToRemove.Close()

	fToKeep, err := os.Open(fileToKeep)
	if err != nil {
		return err
	}
	defer fToKeep.Close()

	scannerToRemove := bufio.NewScanner(fToRemove)
	scannerToKeep := bufio.NewScanner(fToKeep)

	var entryToRemove, entryToKeep FileEntry

	// Initial scan of both files
	hasMoreToRemove := scannerToRemove.Scan()
	if hasMoreToRemove {
		entryToRemove = parseLine(scannerToRemove.Text())
	}

	hasMoreToKeep := scannerToKeep.Scan()
	if hasMoreToKeep {
		entryToKeep = parseLine(scannerToKeep.Text())
	}

	// Iterate through both files, comparing shasums
	for hasMoreToRemove && hasMoreToKeep {
		if entryToRemove.Shasum == entryToKeep.Shasum {

			// repeat until shasums don't match
			for entryToRemove.Shasum == entryToKeep.Shasum {
				if action == "duplicate" {
					fmt.Printf("%s is a duplicate of %s\n", entryToRemove.Path, entryToKeep.Path)
				} else if action == "remove" {
					fmt.Printf("rm \"%s\"\n", entryToRemove.Path)
				}

				totalDuplicates++

				// Move both scanners to the next line
				hasMoreToRemove = scannerToRemove.Scan()
				if hasMoreToRemove {
					entryToRemove = parseLine(scannerToRemove.Text())
				} else {
					// No more lines in fileToRemove, break to avoid infinite loop
					break
				}
			}

			hasMoreToKeep = scannerToKeep.Scan()
			if hasMoreToKeep {
				entryToKeep = parseLine(scannerToKeep.Text())
			}
		} else if entryToRemove.Shasum < entryToKeep.Shasum {
			if action == "unique" {
				fmt.Printf("%s is unique\n", entryToRemove.Path)
			}
			// Move file ToRemove forward
			hasMoreToRemove = scannerToRemove.Scan()
			if hasMoreToRemove {
				entryToRemove = parseLine(scannerToRemove.Text())
			}
		} else {
			if action == "unique" {
				fmt.Printf("%s is unique\n", entryToKeep.Path)
			}

			// Move file ToKeep forward
			hasMoreToKeep = scannerToKeep.Scan()
			if hasMoreToKeep {
				entryToKeep = parseLine(scannerToKeep.Text())
			}
		}
	}

	return nil
}

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Before running the program, you need to generate the shasums of the files to compare using the create_shasums_script.sh script")
		fmt.Printf("Usage: %s <action> <shasums_of_files_to_remove> <shasums_of_files_to_keep>\n", os.Args[0])
		fmt.Println("The action can be 'duplicate', 'unique', 'remove'")
		fmt.Println("duplicate: Find duplicate files")
		fmt.Println("unique: Find unique files")
		fmt.Println("remove: Create script to remove duplicates")
		return
	}

	action := os.Args[1]
	filesToRemove := os.Args[2]
	filesToKeep := os.Args[3]

	if action == "unique" {
		fmt.Println("#Find unique files")
	} else if action == "duplicate" {
		fmt.Println("#Find duplicate files")
	} else if action == "remove" {
		fmt.Println("#Create script to remove duplicates")
	} else {
		fmt.Println("#Invalid action")
		return
	}

	err := compareFiles(filesToRemove, filesToKeep, action)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("#Total duplicates: %d\n", totalDuplicates)
}

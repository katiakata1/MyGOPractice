package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintFilesDirectory(path string) error {
	// You need to open dir fist
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// Loop through directory entries
	for _, entry := range dirEntries {
		if entry.IsDir() {
			fmt.Println("Directory: ", entry.Name())
			newDir := filepath.Join(path, entry.Name())
			if err := PrintFilesDirectory(newDir); err != nil {
				return err
			}
		} else {
			fmt.Println("File: ", entry.Name())
		}
	}
	return nil
}

// func main() {
// 	rootDir := "."

// 	if err := PrintFilesDirectory(rootDir); err != nil {
// 		fmt.Println("Error: ", err)
// 	}
// }

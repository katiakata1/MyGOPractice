package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintFile(path string) error {
	queue := []string{path}

	for len(queue) > 0 {
		dir := queue[0]
		queue = queue[1:]

		entries, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			fullPath := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				fmt.Println("Directory: ", fullPath)
				queue = append(queue, fullPath)
			} else {
				fmt.Println("File: ", fullPath)
			}
		}
	}
	return nil
}

func main() {
	rootDir := "."

	err := PrintFile(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

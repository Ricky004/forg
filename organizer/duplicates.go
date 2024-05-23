package organizer

import (
	"fmt"
	"os"
	"path/filepath"
)

type DuplicateFile struct {
	Original  string
	Duplicate string
}

func DetectDuplicates(dir string) ([]DuplicateFile, error) {
	checksumMap := make(map[string]string)
	duplicates := []DuplicateFile{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		checksum, err := CalculateChecksum(path)
		if err != nil {
			return err
		}

		if original, exists := checksumMap[checksum]; exists {
			duplicates = append(duplicates, DuplicateFile{Original: original, Duplicate: path})
		} else {
			checksumMap[checksum] = path
		}
		return nil
	})

	return duplicates, err
}

func RemoveDuplicates(duplicates []DuplicateFile) error {
	for _, file := range duplicates {
		if err := os.Remove(file.Duplicate); err != nil {
			return err
		}
		fmt.Printf("Removed duplicate file: %s\n", file.Duplicate)
	} 
	return nil
}

func RelocateDuplicates(duplicates []DuplicateFile, targetDir string) error {
	for _, file := range duplicates {
		newPath := filepath.Join(targetDir, filepath.Base(file.Duplicate))
		if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
			return err
		}
		if err := os.Rename(file.Duplicate, newPath); err != nil {
			return err
		}
		fmt.Printf("Moved duplicate file: %s to %s\n", file.Duplicate, newPath)
	}
	return nil
}

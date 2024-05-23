package organizer

import (
	"fmt"
	"os"
	"path/filepath"
)

func OrganizeByDate(dir string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		modtime := info.ModTime()
		year, month, _ := modtime.Date()
		monthStr := fmt.Sprintf("%02d", month)

		targetDir := filepath.Join(dir, fmt.Sprintf("%d", year), monthStr)
		if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
			return err
		}

		newPath := filepath.Join(targetDir, info.Name())
		if err := os.Rename(path, newPath); err != nil {
			return err
		}
		fmt.Printf("Moved %s to %s\n", path, newPath)
		return nil
	})
	return err
}

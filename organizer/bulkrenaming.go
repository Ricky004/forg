package organizer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func BulkRenaming(dir string, prefix, suffix string, startNumber int) error {
   count := startNumber

   err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    if err != nil {
		return nil
	}

	if info.IsDir() {
		return nil
	}

	fileDir := filepath.Dir(path)
	fileExt := filepath.Ext(info.Name())
	fileName := strings.TrimSuffix(info.Name(), fileExt)

	newName := fmt.Sprintf("%s%s%d%s%s", prefix, fileName, count, suffix, fileExt)
	newPath := filepath.Join(fileDir, newName)

	if err := os.Rename(path, newPath); err != nil {
		return err
	}

	fmt.Printf("Renamed %s to %s\n", path, newPath)
	count++
	return nil
	
   })

   return err
}
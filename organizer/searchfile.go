package organizer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func parseSize(sizeStr string) (int64, error) {
	if sizeStr == "" {
		return 0, nil
	}

	units := map[string]int64{
		"b":  1,
        "kb": 1024,
        "mb": 1024 * 1024,
        "gb": 1024 * 1024 * 1024,
	}

	var size int64
	var unit string
	n, err := fmt.Sscanf(sizeStr, "%d%s", &size, &unit)
	if n != 2 || err != nil {
		return 0, fmt.Errorf("invalid size format")
	}

	factor, ok := units[unit]
	if !ok {
		return 0, fmt.Errorf("invalid size unit")
	}

	return size * factor, nil 
}

func SearchFiles(paths []string, searchName, searchExtension, beforeDate, afterDate string, minSizeStr, maxSizeStr string) error {
	var beforeTime, afterTime time.Time
    var err error

	if beforeDate != "" {
		beforeTime, err = time.Parse("2006-01-02", beforeDate)
		if err != nil {
			return fmt.Errorf("invalid before date format")
		}
	}

	if afterDate != "" {
		afterTime, err = time.Parse("2006-01-02", afterDate)
		if err != nil {
			return fmt.Errorf("invalid after date format")
		}
	}

	minSize, err := parseSize(minSizeStr)
    if err != nil {
        return err
    }

    maxSize, err := parseSize(maxSizeStr)
    if err != nil {
        return err
    }

	fileChan := make(chan string, 100)
	var wg sync.WaitGroup
     
	const numWorkers = 10
    for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range fileChan {
				info, err := os.Stat(path)
				if err != nil {
					continue
				}
               
				if !info.Mode().IsRegular() {
					continue
				}
	
				if searchName != "" && !strings.Contains(info.Name(), searchName) {
					continue
				}
	
				if searchExtension != "" && filepath.Ext(info.Name()) != searchExtension {
					continue
				}
	
				if minSize > 0 && info.Size() < minSize {
					continue
				}
				
				if minSize > 0 && info.Size() > maxSize {
					continue
				}
				
				if !beforeTime.IsZero() && info.ModTime().After(beforeTime) {
					continue
				}
	
				if !afterTime.IsZero() && info.ModTime().Before(afterTime) {
					continue
				}
				
				fmt.Println(path)

			}
		}()
	}


	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if os.IsPermission(err) {
					fmt.Printf("Skipping %s: permission denied\n", path)
					return nil
				}
				return err
			}

			if info.Mode().IsRegular() {
				fileChan <- path
			}

			
            return nil
		})

		if err != nil {
			return err
		}
	}

	close(fileChan)
    wg.Wait()

	return nil
}
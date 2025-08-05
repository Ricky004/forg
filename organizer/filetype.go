package organizer

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"github.com/Ricky004/forg/utils"
)

var DefaultCategories = map[string][]string{
	"Images":   {".jpg", ".jpeg", ".png", ".gif"},
	"Documents": {".pdf", ".doc", ".docx", ".txt"},
	"Videos":    {".mp4", ".avi", ".mkv"},
	"Music":     {".mp3", ".wav", ".aac"},
}

type CategoryConfig struct {
	Categories map[string][]string `yaml:"categories"`
}

// load the yaml config file
func LoadConfig(configPath string) (*CategoryConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config CategoryConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func CategorizeByType(dir string, configPath string) error {
	var categories map[string][]string

	if configPath != "" {
		config, err := LoadConfig(configPath)
		if err != nil {
			return err
		}
		categories = config.Categories
	} else {
		categories = DefaultCategories
	}
 
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		for category, extensions := range categories {
			for _, exetension := range extensions {
				if ext == exetension {
					destinationDir := filepath.Join(dir, category)
					if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
						return err
					}

					newPath := filepath.Join(destinationDir, info.Name())
					if err := os.Rename(path, newPath); err != nil {
						return err
					}

					fmt.Printf("Moved %s to %s\n", path, newPath)
					utils.LogOperation(fmt.Sprintf("Moved %s to %s", path, newPath))
					return nil
				}
			}
		} 
		return nil
	})
}

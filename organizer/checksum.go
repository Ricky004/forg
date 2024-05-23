package organizer

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func CalculateChecksum(dir string) (string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
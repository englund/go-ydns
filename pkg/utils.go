package pkg

import (
	"errors"
	"fmt"
	"os"
)

func ReadIpFromFile(filePath string) (string, error) {
	savedIp, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// File does not exist, ignore the error
			return "", nil
		} else {
			return "", fmt.Errorf("error reading IP from file: %w", err)
		}
	}

	return string(savedIp), nil
}

func WriteIpToFile(filePath string, currentIp string) error {
	if err := os.WriteFile(filePath, []byte(currentIp), 0644); err != nil {
		return fmt.Errorf("error writing IP to file: %w", err)
	}

	return nil
}

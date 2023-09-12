package os

import (
	"fmt"
	"os"
)

// ReadFromFile reads data from a file
func ReadFromFile(filePath string, length int64) ([]byte, error) {
	file, err := os.Open(filePath)

	// file not found
	if err != nil {
		return nil, fmt.Errorf("open file error: %s", err)
	}

	defer file.Close()

	// file found, let's try to read it
	data := make([]byte, length)
	count, err := file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("read from file error: %s", err)
	}

	return data[:count], nil
}

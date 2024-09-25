package passwd

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Hello(name string) {
	fmt.Println("Hello ", name)
}

func CheckExists(path string) bool {
	// Check if file exists true/false
	_, error := os.Stat(path)
	return !errors.Is(error, os.ErrNotExist)
}

func CheckCSV(file bool, path string) string {
	// Once the file exists, validate if it has CSV extension
	if file {
		isCsv := strings.HasSuffix(path, "csv")
		if isCsv {
			return "Exists && is CSV"
		} else {
			return "Exists but is not CSV"
		}
	} else {
		return "Not exists"
	}
}

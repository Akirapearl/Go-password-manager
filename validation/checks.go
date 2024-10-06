package validation

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Hello(name string) {
	fmt.Println("Hello ", name)
	fmt.Printf("\n")
}

func FileExists(path string) bool {
	// Checks file exists
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// Returning a bool for valiation checks.
// Function doesn't need to be aware of previous validation.
func IsCSV(path string) bool {
	// Checks file is CSV
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".csv"
}

/*
func CsvStructured(path string) bool {
	// Check CSV is user,account,password
}
*/

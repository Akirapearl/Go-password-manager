package read

import (
	"encoding/csv"
	"fmt"
	"os"
)

func check(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) {
	// Reads file
	// Empty by default
	file, err := os.Open(path)

	// Reads file
	reader := csv.NewReader(file)
	// don't ignore errors.
	records, err := reader.ReadAll()
	check(err)

	fmt.Println(records)

	file.Close()
}

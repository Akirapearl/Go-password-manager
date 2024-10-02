package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"gopassword/validation"
	"os"
)

func run() error {
	fmt.Println("Introduce your password file")

	var path string
	fmt.Scan(&path)

	if !validation.FileExists(path) {
		return errors.New("Validation Failed @ FileExists")
	}

	if !validation.IsCSV(path) {
		return errors.New("Validation Failed @ IsCSV")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	// don't ignore errors.
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	fmt.Println(records)
	return nil

}

func main() {
	validation.Hello("Miriam")

	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

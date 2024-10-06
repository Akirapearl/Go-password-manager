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

	// File variable declaration - Opens file after checking its an existent CSV
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// Reads file
	reader := csv.NewReader(file)
	// don't ignore errors.
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	fmt.Println(records)

	file.Close()

	// Open the CSV file for appending
	fileW, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err != nil {
		return err
	}

	// Create a CSV writer
	writer := csv.NewWriter(fileW)
	defer writer.Flush()

	// Write a new row to the CSV file
	row := []string{"NinjaPop", "Google", "P@ssw0rd"}
	err = writer.Write(row)
	if err != nil {
		panic(err)
	}

	// Reads file
	reader2 := csv.NewReader(fileW)
	// don't ignore errors.
	records2, err := reader2.ReadAll()
	if err != nil {
		return err
	}

	fmt.Println(records2)

	/*---WARNING---*/
	// Every single line of code needs to be above the return nil line below
	// Otherwise it will return unreachable code
	// (function would be already over)
	/*-------------*/
	return nil
}

func main() {
	validation.Hello("Miriam")

	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

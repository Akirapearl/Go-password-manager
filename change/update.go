package change

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvLine struct {
	Column1 string
	Column2 string
	Column3 string
}

var NewValue string

func switchValidate(v int, old *[3]string, new string) {
	switch v {
	case 0:
		old[0] = new
	case 1:
		old[1] = new
	case 2:
		old[2] = new
	default:
		fmt.Println("Invalid field number!")
	}
	NewValue = new // Set the new value globally
}

func UpdateLine(file string) {
	fmt.Print("\033[;1H\033[2J")
	// Get user inputs to identify the record
	var dataUpdate [3]string
	fmt.Println("Update mode - Introduce your username")
	fmt.Scan(&dataUpdate[0])

	fmt.Println("Now, introduce the site or URL for your user (i.e: Google / 'www.google.com')")
	fmt.Scan(&dataUpdate[1])

	fmt.Println("Finally, introduce your password")
	fmt.Scan(&dataUpdate[2])

	// Get the field to update and new value
	var updatePrompt int
	fmt.Println("Which field do you want to change [0 for username, 1 for site, 2 for password]?")
	fmt.Scan(&updatePrompt)

	var newPrompt string
	fmt.Println("Which value would you like to assign to that field?")
	fmt.Scan(&newPrompt)

	// Save a copy of the original data before modifying it
	originalData := dataUpdate

	// Update the selected field in `dataUpdate`
	switchValidate(updatePrompt, &dataUpdate, newPrompt)

	// Read the CSV content
	lines, err := ReadCsv(file)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Track if update happened
	updated := false

	// Iterate through lines and find the matching line
	for i, line := range lines {
		// Check if this line matches the original data criteria
		if line[0] == originalData[0] && line[1] == originalData[1] && line[2] == originalData[2] {
			lines[i][updatePrompt] = NewValue // Apply the new value
			updated = true
			break
		}
	}

	if updated {
		err = WriteCsv(file, lines)
		if err != nil {
			fmt.Println("Error writing CSV:", err)
		} else {
			fmt.Println("Record updated successfully.")
		}
	} else {
		fmt.Println("No matching record found.")
	}
}

// ReadCsv reads the CSV file into a 2D slice
func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// WriteCsv writes the CSV data back to the file
func WriteCsv(filename string, data [][]string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	return writer.WriteAll(data)
}

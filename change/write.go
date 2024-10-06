package change

import (
	"encoding/csv"
	"os"
)

func check(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func WriteLine(file string) {
	// Open the CSV file for appending
	fileW, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	// Create a CSV writer
	writer := csv.NewWriter(fileW)

	// Write a new row to the CSV file
	// Current placed credentials were randomly generated and don't represent a real user
	// Expected structure for the CSV is [user,account/site, password]
	row := []string{"NinjaPop", "Google", "P@ssw0rd"}
	err = writer.Write(row)
	check(err)
	writer.Flush()
}

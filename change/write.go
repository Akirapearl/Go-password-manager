package change

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

func WriteLine(file string) {
	// Ask for values to be added to the file
	fmt.Println("Introduce your username")

	var dataWrite [3]string
	fmt.Scan(&dataWrite[0])

	fmt.Println("Now, introduce the site or URL for your user (i.e: Google / 'www.google.com')")
	fmt.Scan(&dataWrite[1])

	fmt.Println("Finally, introduce your password")
	fmt.Scan(&dataWrite[2])

	// Open the CSV file for appending
	fileW, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	check(err)
	// Create a CSV writer
	writer := csv.NewWriter(fileW)

	// Write a new row to the CSV file
	// Current placed credentials were randomly generated and don't represent a real user
	// Expected structure for the CSV is [user,account/site, password]
	row := []string{dataWrite[0], dataWrite[1], dataWrite[2]}
	err = writer.Write(row)
	check(err)
	writer.Flush()
}

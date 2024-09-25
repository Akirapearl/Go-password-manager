package main

import (
	"encoding/csv"
	"fmt"
	"os"
	passwd "passwd/internal"
)

func main() {
	passwd.Hello("Miriam")

	// Check for existent password CSV
	fmt.Println("Introduce your password file")
	var path string
	fmt.Scan(&path)
	// Send it to external functions - Validate file exists && has CSV extension
	resultPath := passwd.CheckExists(path)

	// If all good - proceed to open it
	if passwd.CheckCSV(resultPath, path) == "Exists && is CSV" {
		// open and read file - intended structure is:
		// user, site/app (account), password
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		reader := csv.NewReader(file)
		records, _ := reader.ReadAll()

		fmt.Println(records)
		// upgrade path: add date of last modification, email, base64 encryption
	} else {
		fmt.Println("Exists but is not CSV")
	}
}

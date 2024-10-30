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

func switchValidate(v int, old *[3]string, new string) (NewValue string) {
	var index int
	switch v {
	case 0:
		old[0] = new
		index = 0
		fmt.Printf("New value is: %v", new)
	case 1:
		old[1] = new
		index = 1
		fmt.Printf("New value is: %v", new)
	case 2:
		old[2] = new
		index = 2
		fmt.Printf("New value is: %v", new)
	default:
		fmt.Println("The number is wrong!")
	}

	NewValue = old[index]

	return NewValue
}
func UpdateLine(file string) {
	// Ask for values to check
	fmt.Println("Update mode - Introduce your username")

	var dataUpdate [3]string
	fmt.Scan(&dataUpdate[0])

	fmt.Println("Now, introduce the site or URL for your user (i.e: Google / 'www.google.com')")
	fmt.Scan(&dataUpdate[1])

	fmt.Println("Finally, introduce your password")
	fmt.Scan(&dataUpdate[2])

	fmt.Println("Which field do you want to change [0/1/2]?", dataUpdate[0], dataUpdate[1], dataUpdate[2])
	var updatePrompt int
	fmt.Scan(&updatePrompt)

	fmt.Println("Which value would you like to assign to that field? ")
	var newPrompt string
	fmt.Scan(&newPrompt)

	//Pass value alongside the content of the array
	switchValidate(updatePrompt, &dataUpdate, newPrompt)

	// Verify the function returns something
	fmt.Println("============================================")

	lines, err := ReadCsv(file)
	check(err)

	// Loop through lines & turn into struct
	for _, line := range lines {
		data := CsvLine{
			Column1: line[0],
			Column2: line[1],
			Column3: line[2],
		}
		if data.Column1 == dataUpdate[updatePrompt] {
			fmt.Println("it worked", data.Column1)
		}
	}

}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

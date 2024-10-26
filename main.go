package main

import (
	"errors"
	"fmt"
	"gopassword/change"
	"gopassword/read"
	"gopassword/validation"
	"time"
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

	// Read content first
	fmt.Println("=== Current content of your file ===")
	read.ReadFile(path)

	time.Sleep(5 * time.Second)
	fmt.Print("\033[;1H\033[2J")
	// Write single line to file
	change.WriteLine(path)

	fmt.Println("=== Updated content of your file - Insert ===")
	read.ReadFile(path)
	/*---WARNING---*/
	// Every single line of code needs to be above the return nil line below
	// Otherwise it will return unreachable code
	// (function would be already over)
	/*-------------*/
	return nil
}

func main() {
	// Custom greeter
	validation.Hello("Miriam")

	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

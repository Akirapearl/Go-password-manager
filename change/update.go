package change

import "fmt"

func UpdateLine(file string) {
	// Ask for values to check
	fmt.Println("Update mode - Introduce your username")

	var dataUpdate [3]string
	fmt.Scan(&dataUpdate[0])

	fmt.Println("Now, introduce the site or URL for your user (i.e: Google / 'www.google.com')")
	fmt.Scan(&dataUpdate[1])

	fmt.Println("Finally, introduce your password")
	fmt.Scan(&dataUpdate[2])

	// Validate if field exists

	fmt.Println("Which field do you want to change?", dataUpdate[0], dataUpdate[1], dataUpdate[2])

}

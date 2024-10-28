package change

import "fmt"

func switchValidate(v int, old *[3]string, new string) string {
	switch v {
	case 0:
		old[0] = new
		fmt.Printf("New value is: %v", new)
	case 1:
		old[1] = new
		fmt.Printf("New value is: %v", new)
	case 2:
		old[2] = new
		fmt.Printf("New value is: %v", new)
	default:
		fmt.Println("The number is wrong!")
	}
	return "none"
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

	// Validate/Verify if field exists
	fmt.Println("Which field do you want to change [0/1/2]?", dataUpdate[0], dataUpdate[1], dataUpdate[2])
	var updatePrompt int
	fmt.Scan(&updatePrompt)

	fmt.Println("Which value would you like to assign to that field? ")
	var newPrompt string
	fmt.Scan(&newPrompt)

	//Pass value alongside the content of the array
	switchValidate(updatePrompt, &dataUpdate, newPrompt)
}

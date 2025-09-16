package main

import (
	Interface "profit-tracker-go/Interfaces"
)

func main() {

	for {
		result := Interface.GetUserInput()
		if result == "4. Exit" {
			break
		} else {
			//TODO: Call and write parser to handle user selection
		}
	}
}

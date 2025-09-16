package Interface

import (
	"github.com/manifoldco/promptui"
)

func GetUserInput() string {
	prompt := promptui.Select{
		Label:    "Select an option",
		Items:    []string{"1. View Data", "2. Insert Data", "3. Settings", "4. Exit"},
		HideHelp: true,
	}

	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return result
}

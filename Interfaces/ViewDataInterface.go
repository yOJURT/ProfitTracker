package Interfaces

import (
	"github.com/manifoldco/promptui"
)

func GetViewDataInput() *promptui.Select {

	prompt := promptui.Select{
		Label:        "Select an option",
		Items:        []string{"1. View Total Win/Loss Amount", "2. View Total Wins", "3. View Total Losses", "4. View Data For Date Range", "5. Prev"},
		HideHelp:     true,
		HideSelected: true,
	}

	return &prompt
}

func GetTestInput() *promptui.Select {
	prompt := promptui.Select{
		Label:        "Select an option",
		Items:        []string{"1. Test", "2. Work", "3. Prev"},
		HideHelp:     true,
		HideSelected: true,
	}

	return &prompt
}

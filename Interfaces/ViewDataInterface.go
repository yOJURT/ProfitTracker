package Interfaces

import (
	"github.com/manifoldco/promptui"
)

func GetViewDataInput() *promptui.Select {

	prompt := promptui.Select{
		Label:        "Select an option",
		Items:        []string{"1. View Total Win/Loss Amount", "2. View Total Wins", "3. View Total Losses", "4. View Data For Date Range", "5. View All Data", "6. Prev"},
		HideHelp:     true,
		HideSelected: true,
	}

	return &prompt
}

func GetViewDataForDateRangeInput() *promptui.Select {
	prompt := promptui.Select{
		Label:        "Select an option",
		Items:        []string{"1. Sigma", "2. Balls", "3. Poopy", "4. Prev"},
		HideHelp:     true,
		HideSelected: true,
	}

	return &prompt
}

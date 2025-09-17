package Interfaces

import (
	"github.com/manifoldco/promptui"
)

func GetMainMenuInput() *promptui.Select {

	prompt := promptui.Select{
		Label:        "Select an option",
		Items:        []string{"1. View Data", "2. Insert Data", "3. Settings", "4. Exit"},
		HideHelp:     true,
		HideSelected: true,
	}

	return &prompt
}

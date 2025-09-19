package Interfaces

import (
	"github.com/manifoldco/promptui"
)

func GetMainMenuInput(items []string) *promptui.Select {

	prompt := promptui.Select{
		Label:        "Select an option",
		Items:        items,
		HideHelp:     true,
		HideSelected: true,
	}

	return &prompt
}

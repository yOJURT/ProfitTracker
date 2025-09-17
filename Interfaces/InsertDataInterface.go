package Interfaces

import (
	"github.com/manifoldco/promptui"
)

func GetInsertDataInput() *promptui.Select {

	prompt := promptui.Select{
		Label:        "Select an option",
		Items:        []string{"1. Insert by Query", "2. Insert by Json Document", "3. Prev"},
		HideHelp:     true,
		HideSelected: true,
	}

	return &prompt
}

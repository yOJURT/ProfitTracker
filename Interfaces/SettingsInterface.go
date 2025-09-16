package Interfaces

import (
	"github.com/manifoldco/promptui"
)

func GetSettingsInput() *promptui.Select {

	prompt := promptui.Select{
		Label:        "Select an option",
		Items:        []string{"1. Set Database Connection String", "2. Set Databse Name", "3. Set Collection", "4. Prev"},
		HideHelp:     true,
		HideSelected: true,
	}

	return &prompt
}

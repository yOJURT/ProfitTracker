package Interfaces

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type Prompt struct {
	prompt *promptui.Select
}

func RunPrompt(prompt string) string {
	if prompt != "Main Menu" {
		prompt = strings.Split(prompt, ". ")[1]
	}

	var PromptToRun Prompt

	switch prompt {
	case "Settings":
		PromptToRun = Prompt{
			prompt: GetSettingsInput(),
		}
	case "Main Menu", "Prev":
		PromptToRun = Prompt{
			prompt: GetMainMenuInput(),
		}
	case "View Data":
		PromptToRun = Prompt{
			prompt: GetViewDataInput(),
		}

	case "Insert Data":
		PromptToRun = Prompt{
			prompt: GetInsertDataInput(),
		}
	}

	_, result, err := PromptToRun.prompt.Run()

	if err != nil {
		panic(err)
	}

	return result
}

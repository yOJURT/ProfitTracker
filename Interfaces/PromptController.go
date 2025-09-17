package Interfaces

import (
	DataStructures "profit-tracker-go/DataStructures"
	"strings"

	"github.com/manifoldco/promptui"
)

type Prompt struct {
	PathStack *DataStructures.Stack
	prompt    *promptui.Select
}

func RunPrompt(prompt string, p *Prompt) string {

	var PromptToRun Prompt
	//TODO: Make this cleaner
	if prompt != "0. Main Menu" {
		prompt = strings.Split(prompt, ". ")[1]
	}

	switch prompt {
	case "0. Main Menu":
		PromptToRun = Prompt{
			prompt: GetMainMenuInput(),
		}
	case "Settings":
		PromptToRun = Prompt{
			prompt: GetSettingsInput(),
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

	if strings.Split(result, ". ")[1] == "Prev" {
		p.PathStack.Pop()
		return p.PathStack.Pop()
	}

	return result
}

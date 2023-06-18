package main 

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor int
}

func initialModel() model {
	return model{
	}
}

func ( m model ) Init() tea.Cmd {
	return nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Go Typer failed to start", err)
		os.Exit(1)
	}
}

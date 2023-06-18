package main 

import (
	"fmt"
	"log"

	"go-typer/modules/words"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type model struct {
	textInput textinput.Model
	err error
}


func initialModel(test string) model {

	ti:= textinput.New()
	ti.Placeholder = test
	ti.Focus()
	ti.Width = 20

	return model{
		textInput: ti,
		err: nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("Go Typer\n\n%s\n\n%s", m.textInput.View(), "(press esc to quit)")
}

func main() {

	test := words.Generate()

	p := tea.NewProgram(initialModel(test))
	if err := p.Start(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		log.Fatal(err)
	}
}

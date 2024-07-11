package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jordanella/teaspoon"
	zone "github.com/lrstanley/bubblezone"
)

//?--------------------------------------------------------------------------------------------------------------------

//* Button Component
/*
 - An example button component struct implementing the Interactable struct
*/
type InteractiveButton struct {
	label       string
	interaction *teaspoon.Interactable
}

//* Creation Method
/*
 - Returns a new button with the Interactable property intialized
 - A ZoneID is defined so mouse events can be attributed to it
 - Click and Hover Handlers are instantiated
 - A custom OnClick method is defined which toggles the state of IsSelected
*/
func NewButton(label string) InteractiveButton {
	return InteractiveButton{
		label: label,
		interaction: &teaspoon.Interactable{
			ID: zone.NewPrefix(),
			Click: &teaspoon.ClickHandler{
				OnClick: func(
					element teaspoon.Interactive, mouseMsg tea.MouseMsg,
				) (teaspoon.Interactive, tea.Cmd) {
					interaction := element.GetInteraction()
					interaction.IsSelected = !interaction.IsSelected
					return element, nil
				}},
			Hover: &teaspoon.HoverHandler{},
		},
	}
}

//* Implement Interactable Interface
/*
 - Satisfy interface definition and provide reference to interaction property
*/
func (b InteractiveButton) GetInteraction() *teaspoon.Interactable {
	return b.interaction
}

//* Button View
/*
 - Handle conditional rendering using IsSelected and IsHovered states
*/
func (b InteractiveButton) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 2)
	if b.interaction.IsSelected {
		style = style.Foreground(lipgloss.Color("#ffffff")).Bold(true)
	} else if b.interaction.IsHovered {
		style = style.Foreground(lipgloss.Color("#44aaaa"))
	}
	return zone.Mark(b.interaction.ID, style.Render(b.label))
}

//?--------------------------------------------------------------------------------------------------------------------

//* Primary Model
/*
 - Define primary model, nesting button components within it
 - Satisfy tea.Model interface definition with Init(), Update(), and View() methods
*/
type MainModel struct {
	buttons []InteractiveButton
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

//* Update Routine
/*
 - Basic KeyMsg handling included for exiting the application
 - Iterates through the buttons, calling the HandleMouseMsg method
 - Updates the button in the model with the handled state
*/
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	case tea.MouseMsg:
		var cmds []tea.Cmd
		for i, button := range m.buttons {
			newButton, cmd := button.interaction.HandleMouseMsg(button, msg)
			m.buttons[i] = newButton.(InteractiveButton)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
		return m, tea.Batch(cmds...)
	}
	return m, nil
}

//* Model View
/*
 - Returns the string output containing the button components
 - Output string is scanned by zone for determination of mouse bounds
*/
func (m MainModel) View() string {
	var buttons []string
	for _, button := range m.buttons {
		buttons = append(buttons, button.View())
	}
	return zone.Scan("Click the buttons:\n\n" + lipgloss.JoinHorizontal(lipgloss.Center, buttons...))
}

//?--------------------------------------------------------------------------------------------------------------------

//* Custom Click Behaviour Example
/*
 - Global manager for zone is instantiated
 - The primary model is declared and three buttons nested within it
 - The bubbletea application is launched using WithAltScreen and WithMouseCellMotion options
*/
func main() {
	zone.NewGlobal()

	model := MainModel{
		buttons: []InteractiveButton{
			NewButton("Button 1"),
			NewButton("Button 2"),
			NewButton("Button 3"),
		},
	}

	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}

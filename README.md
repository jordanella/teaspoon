# TeaSpoon

TeaSpoon is a Go package that provides interactive elements for terminal user interfaces built with [Bubble Tea](https://github.com/charmbracelet/bubbletea). It offers support for clickable, hoverable, draggable, and droppable elements. Enjoy the basic functionality or easily define your own interaction behaviours. The goal of this project is to enable effortless functionality without standing in the way of your creativity. Give it a stir!

## Features

- Click handling (single click, double click, right click)
- Hover detection
- Drag and drop functionality
- Customizable event handlers
- Employs [Bubble Zone](https://github.com/lrstanley/bubblezone) by default

## Installation

```bash
go get github.com/jordanella/teaspoon
```

## Quick Start

Here's a simple example of how to use BubbleInteract:

```go
package main


// Assumes standard tea and zone implementation
// Zone is not required but is employed in default bounds detection
import (
    "github.com/jordanella/teaspoon"
    tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

// Define interactive component with an *Interactable field
type Component struct {
	interaction *teaspoon.Interactable
}

// Add GetInteraction() *Interactable method to return the field
func (c Component) GetInteraction() *teaspoon.Interactable {
	return c.interaction
}

// Example main model containing our component
type MainModel struct{
    tea.Model
    component Component
}

// Initialize as few or many handlers required when creating component
mainModel := MainModel{
    component: Component{
        interaction: &teaspoon.Interactable{
            ID: zone.NewPrefix(),
            Click: &teaspoon.ClickHandler{},
            Hover: &teaspoon.HoverHandler{},
            Drop: &teaspoon.DropHandler{},
            DragEvent: &teaspoon.DragEventHandler{},
        },
    }
}

// Incorporate HandleMouseMsg and/or HandleExternalEvent functions into the Update pipeline
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
        newComponent, cmd := m.component.interaction.HandleMouseMsg(button, msg)
        m.component = newComponent.(Component)
		return m, cmd
	}
	return m, nil
}

// Start configuring your component to render conditionally!
func (c Component) View() string {
    if c.interaction.IsHovered {
        return "Hovered state!"
    }
    return "Normal state"
}
```

## Documentation

For detailed documentation, please see the Go Docs.


## Examples

Check out the examples directory for more detailed usage examples.


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.


## License

This project is licensed under the MIT License - see the LICENSE file for details.
package teaspoon

import (
	tea "github.com/charmbracelet/bubbletea"
)

//?--------------------------------------------------------------------------------------------------------------------

//* Click Behaviour Handler
/*
 - Interface for handling click events and defining local and external event behaviours
 - Default emission behaviour for an event will not occur if a custom behavior is defined
 - Default behaviours do not emit click events unless EmitMessage is set to true
*/
type ClickHandler struct {
	OnClick       func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	OnDoubleClick func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	OnRightClick  func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)

	OnClickEvent       func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	OnDoubleClickEvent func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	OnRightClickEvent  func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)

	EmitMessages bool
}

//?--------------------------------------------------------------------------------------------------------------------

//* Click Events
/*
 - Click event messages to enable responses to external interactions
 - Default behaviours will broadcast if EmitMessages is set to true
*/
type ClickEvent struct {
	ID        string
	EventType ClickEventType
	MouseMsg  tea.MouseMsg
}

//* Click Event Types
/*
 - Enum for providing context to ClickEvent messages
*/
type ClickEventType int

const (
	Click ClickEventType = iota
	DoubleClick
	RightClick
)

//?--------------------------------------------------------------------------------------------------------------------

//* Clickable Interface
/*
 - Interface definition for handling localized click interactions
*/
type Clickable interface {
	HandleClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	HandleDoubleClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	HandleRightClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

//* Left Click Handler
/*
 - Responds to localized clicks with OnClick function or DefaultClick if it is not defined
*/
func (h *ClickHandler) HandleClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnClick != nil {
		return h.OnClick(element, mouseMsg)
	}
	return h.DefaultClick(element, mouseMsg)
}

//* Default Click Behaviour
/*
 - Sets an element's MouseInteraction IsSelected property to true
*/
func (h *ClickHandler) DefaultClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	interaction := element.GetInteraction()
	interaction.IsSelected = true

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return ClickEvent{
				EventType: Click,
				ID:        interaction.ID,
				MouseMsg:  mouseMsg,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Double Left Click Handler
/*
 - Responds to localized double click events with OnDoubleClick function or DefaultDoubleClick if it is not defined
*/
func (h *ClickHandler) HandleDoubleClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnDoubleClick != nil {
		return h.OnDoubleClick(element, mouseMsg)
	}
	return h.DefaultDoubleClick(element, mouseMsg)
}

//* Default Double Click Behaviour
/*
 - Sets an element's MouseInteraction IsSelected property to true
*/
func (h *ClickHandler) DefaultDoubleClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()
	interaction.IsSelected = true

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return ClickEvent{
				EventType: DoubleClick,
				ID:        interaction.ID,
				MouseMsg:  mouseMsg,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Right Click Handler
/*
 - Sets an element's MouseInteraction IsSelected property to true
*/
func (h *ClickHandler) HandleRightClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnRightClick != nil {
		return h.OnRightClick(element, mouseMsg)
	}
	return element, nil
}

//* Default Right Click Behaviour
/*
 - Sets an element's MouseInteraction IsSelected property to true
*/
func (h *ClickHandler) DefaultRightClick(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()
	interaction.IsSelected = true

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return ClickEvent{
				EventType: DoubleClick,
				ID:        interaction.ID,
				MouseMsg:  mouseMsg,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Click Event Aware Interface
/*
 - Interface definition for responding to external click interactions
 - No default behaviour is defined for responding to external click events
*/
type ClickEventAware interface {
	HandleClickEvent(element Interactive, clickEvent ClickEvent) (Interactive, tea.Cmd)
	HandleDoubleClickEvent(element Interactive, clickEvent ClickEvent) (Interactive, tea.Cmd)
	HandleRightClickEvent(element Interactive, clickEvent ClickEvent) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Left Click Handler
/*
 - Responds to external click events with OnClickEvent function if defined
 - No default behaviour is defined for responding to external click events
*/
func (h *ClickHandler) HandleClickEvent(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnClickEvent != nil {
		return h.OnClickEvent(element, mouseMsg)
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Double Left Click Handler
/*
 - Responds to external double click events with OnDoubleClickEvent function if defined
 - No default behaviour is defined for responding to external double click events
*/
func (h *ClickHandler) HandleDoubleClickEvent(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnDoubleClickEvent != nil {
		return h.OnDoubleClickEvent(element, mouseMsg)
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Right Click Handler
/*
 - Responds to external right click events with OnRightClickEvent function if defined
 - No default behaviour is defined for responding to external right click events
*/
func (h *ClickHandler) HandleRightClickEvent(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnRightClickEvent != nil {
		return h.OnRightClickEvent(element, mouseMsg)
	}
	return element, nil
}

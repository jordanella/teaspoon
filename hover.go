package teaspoon

import (
	tea "github.com/charmbracelet/bubbletea"
)

//?--------------------------------------------------------------------------------------------------------------------

//* Hover Behaviour Handler
/*
 - Interface for handling hover events and defining local and external event behaviours
 - Default emission behaviour for an event will not occur if a custom behavior is defined
 - Default behaviours do not emit hover events unless EmitMessage is set to true
*/
type HoverHandler struct {
	OnMouseEnter func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	OnMouseHover func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	OnMouseLeave func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)

	OnMouseEnterEvent func(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd)
	OnMouseHoverEvent func(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd)
	OnMouseLeaveEvent func(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd)

	EmitMessages bool
}

//?--------------------------------------------------------------------------------------------------------------------

//* Hover Events
/*
 - Hover event messages to enable responses to external interactions
 - Default behaviours will broadcast if EmitMessages is set to true
*/
type HoverEvent struct {
	ID        string
	EventType HoverEventType
	MouseMsg  tea.MouseMsg
}

//* Hover Event Types
/*
 - Enum for providing context to HoverEvent messages
*/
type HoverEventType int

const (
	MouseEnter HoverEventType = iota
	MouseHover
	MouseLeave
)

//?--------------------------------------------------------------------------------------------------------------------

//* Hoverable Interface
/*
 - Interface definition for handling localized hover interactions
*/
type Hoverable interface {
	HandleMouseEnter(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	HandleMouseHover(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	HandleMouseLeave(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

//* Mouse Enter Handler
/*
 - Responds to localized mouse enter events with OnMouseEnter function or DefaultMouseEnter if undefined
*/
func (h *HoverHandler) HandleMouseEnter(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnMouseEnter != nil {
		return h.OnMouseEnter(element, mouseMsg)
	}
	return h.DefaultMouseEnter(element, mouseMsg)
}

//* Default Mouse Enter Behaviour
/*
 - Sets an element's MouseInteraction IsHovered property to true
*/
func (h *HoverHandler) DefaultMouseEnter(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	interaction := element.GetInteraction()
	interaction.IsHovered = true

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return HoverEvent{
				EventType: MouseEnter,
				ID:        interaction.ID,
				MouseMsg:  mouseMsg,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Mouse Enter Handler
/*
 - Responds to localized mouse hover events with OnMouseHover function or DefaultMouseHover if undefined
*/
func (h *HoverHandler) HandleMouseHover(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnMouseHover != nil {
		return h.OnMouseHover(element, mouseMsg)
	}
	return h.DefaultMouseHover(element, mouseMsg)
}

//* Default Mouse Enter Behaviour
/*
 - Sets an element's MouseInteraction IsHovered property to true
*/
func (h *HoverHandler) DefaultMouseHover(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	interaction := element.GetInteraction()
	interaction.IsHovered = true

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return HoverEvent{
				EventType: MouseHover,
				ID:        interaction.ID,
				MouseMsg:  mouseMsg,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Mouse Leave Handler
/*
 - Responds to localized mouse leaving events with OnMouseLeave function or DefaultMouseLeave if undefined
*/
func (h *HoverHandler) HandleMouseLeave(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if h.OnMouseLeave != nil {
		return h.OnMouseLeave(element, mouseMsg)
	}
	return h.DefaultMouseLeave(element, mouseMsg)
}

//* Default Mouse Leave Behaviour
/*
 - Sets an element's MouseInteraction IsHovered property to false
*/
func (h *HoverHandler) DefaultMouseLeave(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()
	interaction.IsHovered = false

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return HoverEvent{
				EventType: MouseLeave,
				ID:        interaction.ID,
				MouseMsg:  mouseMsg,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Hover Event Aware Interface
/*
 - Interface definition for handling localized hover interactions
*/
type HoverEventAware interface {
	HandleMouseEnterEvent(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd)
	HandleMouseHoverEvent(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd)
	HandleMouseLeaveEvent(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Mouse Enter Handler
/*
 - Responds to external mouse enter events with OnMouseEnterEvent function if defined
 - No default behaviour is defined for responding to external mouse enter events
*/
func (h *HoverHandler) HandleMouseEnterEvent(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd) {
	if h.OnMouseEnterEvent != nil {
		return h.OnMouseEnterEvent(element, hoverEvent)
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Mouse Hover Handler
/*
 - Responds to external mouse hover events with OnMouseHoverEvent function if defined
 - No default behaviour is defined for responding to external mouse hover events
*/
func (h *HoverHandler) HandleMouseHoverEvent(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd) {
	if h.OnMouseHoverEvent != nil {
		return h.OnMouseHoverEvent(element, hoverEvent)
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Mouse Leave Handler
/*
 - Responds to external mouse leave events with OnMouseLeaveEvent function if defined
 - No default behaviour is defined for responding to external mouse leave events
*/
func (h *HoverHandler) HandleMouseLeaveEvent(element Interactive, hoverEvent HoverEvent) (Interactive, tea.Cmd) {
	if h.OnMouseLeaveEvent != nil {
		return h.OnMouseLeaveEvent(element, hoverEvent)
	}
	return element, nil
}

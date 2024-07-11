package teaspoon

import (
	tea "github.com/charmbracelet/bubbletea"
)

//?--------------------------------------------------------------------------------------------------------------------

//* Drag Behaviour Handler
/*
 - Interface for handling drag events and defining local and external event behaviours
 - Default emission behaviour for an event will not occur if a custom behavior is defined
 - Default behaviours do not emit drag events unless EmitMessage is set to true
*/
type DragHandler struct {
	OnDragStart func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	OnDragMove  func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	OnDragEnd   func(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)

	OnDragStartEvent func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	OnDragMoveEvent  func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	OnDragEndEvent   func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)

	EmitMessages bool
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drag Events
/*
 - Drag event messages to enable responses to external interactions
 - Default behaviours will broadcast if EmitMessages is set to true
*/
type DragEvent struct {
	ID         string
	EventType  DragEventType
	DragType   string
	DragOrigin Point
	DragOffset Point
	MouseMsg   tea.MouseMsg
}

//* Drag Event Types
/*
 - Enum for providing context to DragEvent messages
*/
type DragEventType int

const (
	DragStart DragEventType = iota
	DragMove
	DragEnd
)

//?--------------------------------------------------------------------------------------------------------------------

//* Point Type
/*
 - Helper struct for easily defining a coordinate pair
*/
type Point struct {
	X, Y int
}

//?--------------------------------------------------------------------------------------------------------------------

//* Draggable Interface
/*
 - Interface definition for handling localized drag interactions
*/
type Draggable interface {
	HandleDragStart(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	HandleDragMove(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
	HandleDragEnd(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drag Start Handler
/*
 - Responds to localized drag start events with OnDragStart function or DefaultDragStart if undefined
*/
func (h *DragHandler) HandleDragStart(element Interactive, mouseMsg tea.MouseMsg, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDragStart != nil {
		return h.OnDragStart(element, mouseMsg)
	}
	return h.DefaultDragStart(element, mouseMsg)
}

//* Default Drag Start Behaviour
/*
 - Sets an element's MouseInteraction IsDragging property to true
*/
func (h *DragHandler) DefaultDragStart(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	interaction := element.GetInteraction()

	interaction.IsDragging = true
	interaction.DragOrigin = Point{X: mouseMsg.X, Y: mouseMsg.Y}
	interaction.DragOffset = Point{X: 0, Y: 0}

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return DragEvent{
				EventType:  DragStart,
				ID:         interaction.ID,
				MouseMsg:   mouseMsg,
				DragOrigin: interaction.DragOrigin,
				DragOffset: interaction.DragOffset,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drag Move Handler
/*
 - Responds to localized drag move events with OnDragMove function or DefaultDragMove if undefined
*/
func (h *DragHandler) HandleDragMove(element Interactive, mouseMsg tea.MouseMsg, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDragMove != nil {
		return h.OnDragMove(element, mouseMsg)
	}
	return h.DefaultDragMove(element, mouseMsg)
}

//* Default Drag Move Behaviour
/*
 - Sets an element's MouseInteraction IsDragging property to true
*/
func (h *DragHandler) DefaultDragMove(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	interaction := element.GetInteraction()

	interaction.IsDragging = true
	interaction.DragOffset = Point{
		X: mouseMsg.X - interaction.DragOrigin.X,
		Y: mouseMsg.Y - interaction.DragOrigin.Y,
	}

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return DragEvent{
				EventType:  DragMove,
				ID:         interaction.ID,
				MouseMsg:   mouseMsg,
				DragOrigin: interaction.DragOrigin,
				DragOffset: interaction.DragOffset,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drag End Handler
/*
 - Responds to localized drag end events with OnDragEnd function or DefaultDragEnd if undefined
*/
func (h *DragHandler) HandleDragEnd(element Interactive, mouseMsg tea.MouseMsg, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDragEnd != nil {
		return h.OnDragEnd(element, mouseMsg)
	}
	return h.DefaultDragEnd(element, mouseMsg)
}

//* Default Drag End Behaviour
/*
- Sets an element's MouseInteraction IsDragging property to false
*/
func (h *DragHandler) DefaultDragEnd(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	interaction := element.GetInteraction()
	interaction.IsDragging = false

	if h.EmitMessages {
		cmd := func() tea.Msg {
			return DragEvent{
				EventType:  DragEnd,
				ID:         interaction.ID,
				MouseMsg:   mouseMsg,
				DragOrigin: interaction.DragOrigin,
				DragOffset: interaction.DragOffset,
			}
		}
		return element, cmd
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drag Event Aware Interface
/*
 - Interface definition for handling localized hover interactions
*/
type DragEventAware interface {
	HandleDragStartEvent(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	HandleDragMoveEvent(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	HandleDragEndEvent(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drag Start Handler
/*
 - Responds to external drag start events with OnDragStartEvent function if defined
 - No default behaviour is defined for responding to external drag start events
*/
func (h *DragHandler) HandleDragStartEvent(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDragStartEvent != nil {
		return h.OnDragStartEvent(element, dragEvent)
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drag Move Handler
/*
 - Responds to external mouse leave events with OnDragMoveEvent function if defined
 - No default behaviour is defined for responding to external drag move events
*/
func (h *DragHandler) HandleDragMoveEvent(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDragMoveEvent != nil {
		return h.OnDragMoveEvent(element, dragEvent)
	}
	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drag End Handler
/*
 - Responds to external drag end events with OnDragEndEvent function if defined
 - No default behaviour is defined for responding to external drag end events
*/
func (h *DragHandler) HandleDragEndEvent(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDragEndEvent != nil {
		return h.OnDragEndEvent(element, dragEvent)
	}
	return element, nil
}

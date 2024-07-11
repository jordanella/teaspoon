package teaspoon

import (
	tea "github.com/charmbracelet/bubbletea"
)

//?--------------------------------------------------------------------------------------------------------------------

//* Drop Behaviour Handler
/*
 - Interface for handling drop events and defining local and external event behaviours
 - Default emission behaviour for an event will not occur if a custom behavior is defined
 - Default behaviours do not emit hover events unless EmitMessage is set to true
*/
type DropHandler struct {
	OnDropEnter   func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	OnDropHover   func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	OnDropLeave   func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	OnDropRelease func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	OnDropAccept  func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	OnDropDeny    func(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	IsAcceptable  func(element Interactive, dragEvent DragEvent) bool

	OnDropEnterEvent   func(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	OnDropHoverEvent   func(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	OnDropLeaveEvent   func(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	OnDropReleaseEvent func(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	OnDropAcceptEvent  func(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	OnDropDenyEvent    func(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)

	AcceptedDropTypes []string
	EmitMessages      bool
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drop Events
/*
 - Drop event messages to enable responses to external interactions
 - Default behaviours will broadcast if EmitMessages is set to true
*/
type DropEvent struct {
	ID         string
	EventType  DropEventType
	DropType   string
	Acceptable bool
	DragEvent  DragEvent
}

//* Drop Event Types
/*
 - Enum for providing context to DropEvent messages
*/
type DropEventType int

const (
	DropEnter DropEventType = iota
	DropHover
	DropLeave
	DropRelease
	DropAccept
	DropDeny
)

//?--------------------------------------------------------------------------------------------------------------------

//* Droppable Interface
/*
 - Interface definition for handling localized drop interactions
*/
type Droppable interface {
	HandleIsAcceptable(element Interactive, dragEvent DragEvent) bool

	HandleDropEnter(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	HandleDropHover(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	HandleDropLeave(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	HandleDropRelease(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	HandleDropAccept(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
	HandleDropDeny(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

//* Acceptable Drop Assessment Handler
/*
 - Responds to acceptable drop assessments with IsAcceptable function or DefaultIsAcceptable if undefined
*/
func (h *DropHandler) HandleIsAcceptable(element Interactive, dragEvent DragEvent) bool {
	if h.IsAcceptable != nil {
		return h.IsAcceptable(element, dragEvent)
	}
	return h.DefaultIsAcceptable(element, dragEvent)
}

//* Default Acceptable Drop Assessment
/*
 - Returns true if the DragType is found within the accepted drop types list
*/
func (h *DropHandler) DefaultIsAcceptable(element Interactive, dragEvent DragEvent) bool {
	for _, dropType := range h.AcceptedDropTypes {
		if dragEvent.DragType == dropType {
			return true
		}
	}
	return false
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drop Enter Handler
/*
 - Responds to localized drop enter events with OnDropEnter function or DefaultDropEnter if undefined
*/
func (h *DropHandler) HandleDropEnter(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDropEnter != nil {
		return h.OnDropEnter(element, dragEvent)
	}
	return h.DefaultDropEnter(element, dragEvent)
}

//* Default Drop Enter
/*
 - Sets an element's MouseInteraction IsBelowDrop property to true and IsValidDrop is determined
*/
func (h *DropHandler) DefaultDropEnter(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	var cmd tea.Cmd

	interaction := element.GetInteraction()
	interaction.IsBelowDrop = true
	interaction.IsValidDrop = h.HandleIsAcceptable(element, dragEvent)

	if h.EmitMessages {
		cmd = func() tea.Msg {
			return DropEvent{
				EventType:  DropEnter,
				ID:         interaction.ID,
				Acceptable: interaction.IsValidDrop,
				DragEvent:  dragEvent,
			}
		}
	}
	return element, cmd
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drop Hover Handler
/*
 - Responds to localized drop hover events with OnDropHover function or DefaultDropHover if undefined
*/
func (h *DropHandler) HandleDropHover(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDropHover != nil {
		return h.OnDropHover(element, dragEvent)
	}
	return h.DefaultDropHover(element, dragEvent)
}

//* Default Drop Hover
/*
 - Sets an element's MouseInteraction IsBelowDrop property to true and IsValidDrop is determined
*/
func (h *DropHandler) DefaultDropHover(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	var cmd tea.Cmd

	interaction := element.GetInteraction()
	interaction.IsBelowDrop = true
	interaction.IsValidDrop = h.HandleIsAcceptable(element, dragEvent)

	if h.EmitMessages {
		cmd = func() tea.Msg {
			return DropEvent{
				EventType:  DropHover,
				ID:         interaction.ID,
				Acceptable: interaction.IsValidDrop,
				DragEvent:  dragEvent,
			}
		}
	}
	return element, cmd
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drop Leave Handler
/*
 - Responds to localized drop leave events with OnDropLeave function or DefaultDropLeave if undefined
*/
func (h *DropHandler) HandleDropLeave(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDropLeave != nil {
		return h.OnDropLeave(element, dragEvent)
	}
	return h.DefaultDropLeave(element, dragEvent)
}

//* Default Drop Leave
/*
 - Sets an element's MouseInteraction IsBelowDrop property to false and IsValidDrop to false
*/
func (h *DropHandler) DefaultDropLeave(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	var cmd tea.Cmd

	interaction := element.GetInteraction()
	interaction.IsBelowDrop = false

	if h.EmitMessages {
		cmd = func() tea.Msg {
			return DropEvent{
				EventType:  DropLeave,
				ID:         interaction.ID,
				Acceptable: interaction.IsValidDrop,
				DragEvent:  dragEvent,
			}
		}
	}

	interaction.IsValidDrop = false

	return element, cmd
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drop Release Handler
/*
 - Responds to localized drop release events with OnDropRelease function or DefaultDropRelease if undefined
*/
func (h *DropHandler) HandleDropRelease(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDropRelease != nil {
		return h.OnDropRelease(element, dragEvent)
	}
	return h.DefaultDropRelease(element, dragEvent)
}

//* Default Drop Release
/*
 - The element's MouseInteraction IsValidDrop is determined
 - Calls relevant accept or deny drop handler method
*/
func (h *DropHandler) DefaultDropRelease(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	interaction := element.GetInteraction()
	interaction.IsValidDrop = h.HandleIsAcceptable(element, dragEvent)

	if h.EmitMessages {
		cmds = append(cmds, func() tea.Msg {
			return DropEvent{
				EventType:  DropRelease,
				ID:         interaction.ID,
				Acceptable: interaction.IsValidDrop,
				DragEvent:  dragEvent,
			}
		})
	}

	if interaction.IsValidDrop {
		element, cmd = h.HandleDropAccept(element, dragEvent)
	} else {
		element, cmd = h.HandleDropDeny(element, dragEvent)
	}

	return element, tea.Batch(append(cmds, cmd)...)
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drop Accept Handler
/*
 - Responds to localized drop accept events with OnDropAccept function or DefaultDropAccept if undefined
*/
func (h *DropHandler) HandleDropAccept(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDropAccept != nil {
		return h.OnDropAccept(element, dragEvent)
	}
	return h.DefaultDropAccept(element, dragEvent)
}

//* Default Drop Accept
/*
 - Sets an element's MouseInteraction IsBelowDrop property to false and IsValidDrop to false
*/
func (h *DropHandler) DefaultDropAccept(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	var cmd tea.Cmd

	interaction := element.GetInteraction()

	if h.EmitMessages {
		cmd = func() tea.Msg {
			return DropEvent{
				EventType:  DropAccept,
				ID:         interaction.ID,
				Acceptable: true,
				DragEvent:  dragEvent,
			}
		}
	}

	interaction.IsBelowDrop = false
	interaction.IsValidDrop = false

	return element, cmd
}

//?--------------------------------------------------------------------------------------------------------------------

//* Drop Deny Handler
/*
 - Responds to localized drop deny events with OnDropDeny function or DefaultDropDeny if undefined
*/
func (h *DropHandler) HandleDropDeny(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	if h.OnDropDeny != nil {
		return h.OnDropDeny(element, dragEvent)
	}
	return h.DefaultDropDeny(element, dragEvent)
}

//* Default Drop Deny
/*
 - Sets an element's MouseInteraction IsBelowDrop property to false and IsValidDrop to false
*/
func (h *DropHandler) DefaultDropDeny(element Interactive, dragEvent DragEvent) (Interactive, tea.Cmd) {
	var cmd tea.Cmd

	interaction := element.GetInteraction()

	if h.EmitMessages {
		cmd = func() tea.Msg {
			return DropEvent{
				EventType:  DropDeny,
				ID:         interaction.ID,
				Acceptable: false,
				DragEvent:  dragEvent,
			}
		}
	}

	interaction.IsBelowDrop = false
	interaction.IsValidDrop = false

	return element, cmd
}

//?--------------------------------------------------------------------------------------------------------------------

//* Droppable Interface
/*
 - Interface definition for handling external drop interactions
*/
type DropEventAware interface {
	HandleDropEnterEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	HandleDropHoverEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	HandleDropLeaveEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	HandleDropReleaseEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	HandleDropAcceptEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
	HandleDropDenyEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drop Enter Handler
/*
 - Responds to external drop enter events with OnDropEnterEvent function if defined
*/
func (h *DropHandler) HandleDropEnterEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {
	if h.OnDropEnterEvent != nil {
		return h.OnDropEnterEvent(element, dropEvent)
	}
	return element, nil
}

//* Default External Drop Enter
/*
 - Sets an element's MouseInteraction IsAboveDrop property to true and IsValidDrop appropriately
*/
func (h *DropHandler) DefaultDropEnterEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()

	interaction.IsAboveDrop = true
	interaction.IsValidDrop = dropEvent.Acceptable

	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drop Hover Handler
/*
 - Responds to external drop hover events with OnDropHoverEvent function if defined
*/
func (h *DropHandler) HandleDropHoverEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {
	if h.OnDropEnterEvent != nil {
		return h.OnDropHoverEvent(element, dropEvent)
	}
	return h.DefaultDropHoverEvent(element, dropEvent)
}

//* Default External Drop Hover
/*
 - Sets an element's MouseInteraction IsAboveDrop property to true and IsValidDrop appropriately
*/
func (h *DropHandler) DefaultDropHoverEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()

	interaction.IsAboveDrop = true
	interaction.IsValidDrop = dropEvent.Acceptable

	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drop Leave Handler
/*
 - Responds to external drop leave events with OnDropLeaveEvent function if defined
*/
func (h *DropHandler) HandleDropLeaveEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {
	if h.OnDropLeaveEvent != nil {
		return h.OnDropLeaveEvent(element, dropEvent)
	}
	return h.DefaultDropLeaveEvent(element, dropEvent)
}

//* Default External Drop Leave
/*
 - Sets an element's MouseInteraction IsAboveDrop property to false and IsValidDrop appropriately
*/
func (h *DropHandler) DefaultDropLeaveEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()

	interaction.IsAboveDrop = false
	interaction.IsValidDrop = dropEvent.Acceptable

	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drop Release Handler
/*
 - Responds to external drop release events with OnDropReleaseEvent function if defined
*/
func (h *DropHandler) HandleDropReleaseEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {
	if h.OnDropReleaseEvent != nil {
		return h.OnDropReleaseEvent(element, dropEvent)
	}
	return h.DefaultDropReleaseEvent(element, dropEvent)
}

//* Default Drop Release
/*
 - Sets an element's MouseInteraction IsAboveDrop property to false and IsValidDrop to false
*/
func (h *DropHandler) DefaultDropReleaseEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()
	interaction.IsAboveDrop = false
	interaction.IsValidDrop = false

	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drop Accept Handler
/*
 - Responds to external drop accept events with OnDropAcceptEvent function if defined
*/
func (h *DropHandler) HandleDropAcceptEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {
	if h.OnDropAcceptEvent != nil {
		return h.OnDropAcceptEvent(element, dropEvent)
	}
	return h.DefaultDropAcceptEvent(element, dropEvent)
}

//* Default Drop Accept
/*
 - Sets an element's MouseInteraction IsAboveDrop property to false and IsValidDrop to false
*/
func (h *DropHandler) DefaultDropAcceptEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()
	interaction.IsAboveDrop = false
	interaction.IsValidDrop = false

	return element, nil
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Drop Deny Handler
/*
 - Responds to external drop deny events with OnDropDenyEvent function if defined
*/
func (h *DropHandler) HandleDropDenyEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {
	if h.OnDropDenyEvent != nil {
		return h.OnDropDenyEvent(element, dropEvent)
	}
	return h.DefaultDropDenyEvent(element, dropEvent)
}

//* Default Drop Deny
/*
 - Sets an element's MouseInteraction IsAboveDrop property to false and IsValidDrop to false
*/
func (h *DropHandler) DefaultDropDenyEvent(element Interactive, dropEvent DropEvent) (Interactive, tea.Cmd) {

	interaction := element.GetInteraction()
	interaction.IsAboveDrop = false
	interaction.IsValidDrop = false

	return element, nil
}

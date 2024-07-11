package teaspoon

import (
	tea "github.com/charmbracelet/bubbletea"
)

//?--------------------------------------------------------------------------------------------------------------------

//* Mouse Message Handling
/*
 - Responds to mouse messages with LocalHandler function or DefaultLocalHandler if it is not defined
*/
func (i *Interactable) HandleMouseMsg(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if i.LocalHandler != nil {
		return i.LocalHandler(element, mouseMsg)
	}
	return i.DefaultLocalHandler(element, mouseMsg)
}

//* Default Local Handler
/*
 - Default implementation of interpretting mouse messages to direct the appropriate interaction handlers
*/
func (i *Interactable) DefaultLocalHandler(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {

	if i.Click == nil && i.Hover == nil && i.Drag == nil && i.Drop == nil {
		return element, nil
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd

	var isInside bool
	if i.IsInside != nil {
		isInside = i.IsInside(element, mouseMsg)
	} else {
		isInside = i.DefaultIsInside(element, mouseMsg)
	}

	switch mouseMsg.Action {
	case tea.MouseActionMotion:
		if i.Hover != nil {
			if isInside {
				if !i.IsHovered {
					// Mouse Enter
					element, cmd = i.Hover.HandleMouseEnter(element, mouseMsg)
					if cmd != nil {
						cmds = append(cmds, cmd)
					}
				}
				// Mouse Hover
				element, cmd = i.Hover.HandleMouseHover(element, mouseMsg)
				cmds = append(cmds, cmd)
			} else if i.IsHovered {
				// Mouse Leave
				element, cmd = i.Hover.HandleMouseLeave(element, mouseMsg)
				if cmd != nil {
					cmds = append(cmds, cmd)
				}
			}
		}

		if i.Drag != nil && i.IsDragging {
			// Drag Move
			element, cmd = i.Drag.HandleDragMove(element, mouseMsg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case tea.MouseActionPress:
		if i.Click != nil && isInside {
			element, cmd = i.Click.HandleClick(element, mouseMsg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}

			if i.Drag != nil {
				i.IsDragging = true
				element, cmd = i.Drag.HandleDragStart(element, mouseMsg)
				if cmd != nil {
					cmds = append(cmds, cmd)
				}
			}
		}

	case tea.MouseActionRelease:
		if i.Drag != nil && i.IsDragging {
			i.IsDragging = false
			element, cmd = i.Drag.HandleDragEnd(element, mouseMsg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}

	interaction := element.GetInteraction()
	*interaction = *i

	return element, tea.Batch(cmds...)
}

//?--------------------------------------------------------------------------------------------------------------------

//* External Event Handling
/*
 - Responds to external event messages with ExternalHandler function or DefaultExternalHandler if it is not defined
*/
func (i *Interactable) HandleExternalEvent(element Interactive, mouseMsg tea.MouseMsg) (Interactive, tea.Cmd) {
	if i.LocalHandler != nil {
		return i.ExternalHandler(element, mouseMsg)
	}
	return i.DefaultExternalHandler(element, mouseMsg)
}

//* External Event Handling
/*
 - Interprets external event messages to direct the appropriate interaction handlers
*/
func (i *Interactable) DefaultExternalHandler(element Interactive, msg tea.Msg) (Interactive, tea.Cmd) {

	if i.ClickEvent == nil && i.HoverEvent == nil && i.DragEvent == nil && i.DropEvent == nil {
		return element, nil
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case ClickEvent:
		if i.ClickEvent != nil {
			switch msg.EventType {
			case Click:
				element, cmd = i.ClickEvent.HandleClickEvent(element, msg)
			case DoubleClick:
				element, cmd = i.ClickEvent.HandleDoubleClickEvent(element, msg)
			case RightClick:
				element, cmd = i.ClickEvent.HandleRightClickEvent(element, msg)
			}
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case HoverEvent:
		if i.HoverEvent != nil {
			switch msg.EventType {
			case MouseEnter:
				element, cmd = i.HoverEvent.HandleMouseEnterEvent(element, msg)
			case MouseHover:
				element, cmd = i.HoverEvent.HandleMouseHoverEvent(element, msg)
			case MouseLeave:
				element, cmd = i.HoverEvent.HandleMouseLeaveEvent(element, msg)
			}
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case DragEvent:
		if i.DragEvent != nil {
			switch msg.EventType {
			case DragStart:
				element, cmd = i.DragEvent.HandleDragStartEvent(element, msg)
			case DragMove:
				element, cmd = i.DragEvent.HandleDragMoveEvent(element, msg)
			case DragEnd:
				element, cmd = i.DragEvent.HandleDragEndEvent(element, msg)
			}
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case DropEvent:
		if i.DropEvent != nil {
			switch msg.EventType {
			case DropEnter:
				element, cmd = i.DropEvent.HandleDropEnterEvent(element, msg)
			case DropHover:
				element, cmd = i.DropEvent.HandleDropHoverEvent(element, msg)
			case DropLeave:
				element, cmd = i.DropEvent.HandleDropLeaveEvent(element, msg)
			case DropRelease:
				element, cmd = i.DropEvent.HandleDropReleaseEvent(element, msg)
			case DropAccept:
				element, cmd = i.DropEvent.HandleDropAcceptEvent(element, msg)
			case DropDeny:
				element, cmd = i.DropEvent.HandleDropDenyEvent(element, msg)
			}
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}

	return element, tea.Batch(cmds...)
}

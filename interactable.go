package teaspoon

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

//?--------------------------------------------------------------------------------------------------------------------

//* Interactive Interface
/*
 - Interface definition for defining an element as interactive
*/
type Interactive interface {
	GetInteraction() *Interactable
}

//?--------------------------------------------------------------------------------------------------------------------

//* Mouse Interaction
/*
 -  Defines and handles mouse interaction of and between elements.
*/
type Interactable struct {
	ID string

	LastClickTime        time.Time
	DoubleClickThreshold time.Duration
	ClickCount           int
	IsSelected           bool
	IsHovered            bool
	IsDragging           bool
	DragOrigin           struct{ X, Y int }
	DragOffset           struct{ X, Y int }
	IsValidDrop          bool
	IsAboveDrop          bool
	IsBelowDrop          bool

	Click Clickable
	Hover Hoverable
	Drag  Draggable
	Drop  Droppable

	ClickEvent ClickEventAware
	HoverEvent HoverEventAware
	DragEvent  DragEventAware
	DropEvent  DropEventAware

	IsInside        func(element Interactive, mouseMsg tea.Msg) bool
	LocalHandler    func(element Interactive, msg tea.Msg) (Interactive, tea.Cmd)
	ExternalHandler func(element Interactive, msg tea.Msg) (Interactive, tea.Cmd)
}

//?--------------------------------------------------------------------------------------------------------------------

func (i Interactable) DefaultIsInside(element Interactive, mouseMsg tea.MouseMsg) bool {
	return zone.Get(i.ID).InBounds(mouseMsg)
}

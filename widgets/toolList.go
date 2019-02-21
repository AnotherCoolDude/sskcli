package widgets

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

// ToolList displays a list of tools
type ToolList struct {
	list   *widgets.List
	tools  []string
	active bool
}

// NewToolList returns a new ToolList
func NewToolList(tools []string) *ToolList {
	tl := ToolList{
		list:   widgets.NewList(),
		tools:  tools,
		active: false,
	}
	tl.list.Rows = tools
	tl.list.SelectedRow = 0
	return &tl
}

// SetTitle sets the title
func (tl *ToolList) SetTitle(title string) {
	tl.list.Title = title
}

// SelectedRow returns the index of the selected row
func (tl *ToolList) SelectedRow() uint {
	return tl.list.SelectedRow
}

// satisfy navigatable
func (tl *ToolList) up() {
	tl.list.ScrollUp()
}

func (tl *ToolList) down() {
	tl.list.ScrollDown()
}

func (tl *ToolList) setActive(active bool) {
	tl.active = active
	if active {
		tl.list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack)
	} else {
		tl.list.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)
	}
}

// SelectedRowContent returns the content of the selected row
func (tl *ToolList) SelectedRowContent() string {
	return tl.list.Rows[tl.SelectedRowIndex()]
}

// SelectedRowIndex returns the selected row index
func (tl *ToolList) SelectedRowIndex() uint {
	return tl.list.SelectedRow
}

// Griditem returns the drawable interface needed for ui.grid
func (tl *ToolList) Griditem() ui.Drawable {
	return tl.list
}

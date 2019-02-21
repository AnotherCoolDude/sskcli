package widgets

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

// DependencyList shows requirements based on a given map key
type DependencyList struct {
	list         *widgets.List
	requirements map[string][]string
	active       bool
}

// NewDependencyList returns a new DependencyList
func NewDependencyList(dependencies map[string][]string) *DependencyList {
	dl := DependencyList{
		list:         widgets.NewList(),
		requirements: dependencies,
		active:       false,
	}
	dl.list.SelectedRow = 0
	return &dl
}

// ListDependencies lists the values of the provided key
func (dl *DependencyList) ListDependencies(key string) {
	dl.list.Rows = dl.requirements[key]
}

// satisfy navigatable
func (dl *DependencyList) up() {
	dl.list.ScrollUp()
}

func (dl *DependencyList) down() {
	dl.list.ScrollDown()
}

func (dl *DependencyList) setActive(active bool) {
	dl.active = active
	if active {
		dl.list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack)
	} else {
		dl.list.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)
	}
}

// SelectedRowContent returns the content of the selected row
func (dl *DependencyList) SelectedRowContent() string {
	return dl.list.Rows[dl.SelectedRowIndex()]
}

// SelectedRowIndex returns the selected row index
func (dl *DependencyList) SelectedRowIndex() uint {
	return dl.list.SelectedRow
}

// Griditem returns the drawable interface needed for ui.grid
func (dl *DependencyList) Griditem() ui.Drawable {
	return dl.list
}

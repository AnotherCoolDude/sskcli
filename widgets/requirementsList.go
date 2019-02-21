package widgets

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

// RequirementsList shows requirements based on a given map key
type RequirementsList struct {
	list         *widgets.List
	requirements map[string][]string
	active       bool
}

// NewRequirementsList returns a new RequirementsList
func NewRequirementsList(requirements map[string][]string) *RequirementsList {
	rl := RequirementsList{
		list:         widgets.NewList(),
		requirements: requirements,
		active:       false,
	}
	rl.list.SelectedRow = 0
	return &rl
}

// ListRequirements lists the values of the provided key
func (rl *RequirementsList) ListRequirements(key string) {
	rl.list.Rows = rl.requirements[key]
}

// satisfy navigatable
func (rl *RequirementsList) up() {
	rl.list.ScrollUp()
}

func (rl *RequirementsList) down() {
	rl.list.ScrollDown()
}

func (rl *RequirementsList) setActive(active bool) {
	rl.active = active
	if active {
		rl.list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack)
	} else {
		rl.list.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)
	}
}

// SelectedRowContent returns the content of the selected row
func (rl *RequirementsList) SelectedRowContent() string {
	return rl.list.Rows[rl.SelectedRowIndex()]
}

// SelectedRowIndex returns the selected row index
func (rl *RequirementsList) SelectedRowIndex() uint {
	return rl.list.SelectedRow
}

// Griditem returns the drawable interface needed for ui.grid
func (rl *RequirementsList) Griditem() ui.Drawable {
	return rl.list
}

package widgets

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

type RequirementsList struct {
	list         *widgets.List
	requirements map[string][]string
	active       bool
}

func newRequirementsList(requirements map[string][]string) *requirementsList {
	rl := requirementsList{
		list:         widgets.NewList(),
		requirements: requirements,
		active:       false,
	}
	return &rl
}

func (rl *requirementsList) listRequirements(key string) {
	rl.list.Rows = rl.requirements[key]
}

// satisfy navigatable
func (rl *requirementsList) up() {
	rl.list.ScrollUp()
}

func (rl *requirementsList) down() {
	rl.list.ScrollDown()
}

func (rl *requirementsList) setActive(active bool) {
	rl.active = active
	if active {
		rl.list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack)
	} else {
		rl.list.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)
	}
}

func (rl *requirementsList) selectedRowContent() string {
	return rl.list.Rows[rl.selectedRowIndex()]
}

func (rl *requirementsList) selectedRowIndex() uint {
	return rl.list.SelectedRow
}

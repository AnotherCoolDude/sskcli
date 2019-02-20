package widgets

import ui "github.com/gizak/termui"

type navigator struct {
	items *[]navigatable
	index int
	grid  *ui.Grid
}

func newNavigator(items *[]navigatable, grid *ui.Grid) *navigator {
	return &navigator{
		items: items,
		index: 0,
		grid:  grid,
	}
}

func (nav *navigator) focusOnNextItem() {
	//fmt.Printf("index + 1 = %d, len items = %d\n", nav.index+1, len(*nav.items))
	(*nav.items)[nav.index].setActive(false)
	nav.index++
	if nav.index >= len(*nav.items) {
		nav.index = 0
	}
	(*nav.items)[nav.index].setActive(true)
	ui.Render(nav.grid)
}

func (nav *navigator) up() {
	(*nav.items)[nav.index].up()
	ui.Render(nav.grid)
}

func (nav *navigator) down() {
	(*nav.items)[nav.index].down()
	ui.Render(nav.grid)
}

type navigatable interface {
	setActive(active bool)
	up()
	down()
	selectedRowContent() string
	selectedRowIndex() uint
}

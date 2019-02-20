package widgets

import ui "github.com/gizak/termui"

// Navigator enables tab-based navigation through privded navigatable items
type Navigator struct {
	items *[]Navigatable
	index int
	grid  *ui.Grid
}

// NewNavigator returns a new Navigator struct
func NewNavigator(items *[]Navigatable, grid *ui.Grid) *Navigator {
	return &Navigator{
		items: items,
		index: 0,
		grid:  grid,
	}
}

// FocusOnNextItem changes focus to the next item in items
func (nav *Navigator) FocusOnNextItem() {
	//fmt.Printf("index + 1 = %d, len items = %d\n", nav.index+1, len(*nav.items))
	(*nav.items)[nav.index].setActive(false)
	nav.index++
	if nav.index >= len(*nav.items) {
		nav.index = 0
	}
	(*nav.items)[nav.index].setActive(true)
	nav.RenderItems()
}

// Up scrolls one row up in the focused item
func (nav *Navigator) Up() {
	(*nav.items)[nav.index].up()
	nav.RenderItems()
}

// Down scrolls one row down in the focused item
func (nav *Navigator) Down() {
	(*nav.items)[nav.index].down()
	nav.RenderItems()
}

//FocusedItem returns the currently focused item
func (nav *Navigator) FocusedItem() Navigatable {
	return (*nav.items)[nav.index]
}

// Navigatable satisfies the navigatable interface
type Navigatable interface {
	setActive(active bool)
	up()
	down()
	SelectedRowContent() string
	SelectedRowIndex() uint
	Griditem() ui.Drawable
}

// RenderItems renders all navigator items
func (nav *Navigator) RenderItems() {
	for _, item := range *nav.items {
		ui.Render(item.Griditem())
	}
}

// FocusOnItem focuses on the item at index index
func (nav *Navigator) FocusOnItem(index int) {
	nav.index = index
	nav.RenderItems()
}

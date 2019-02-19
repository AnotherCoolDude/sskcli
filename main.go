// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gizak/termui/widgets"

	ui "github.com/gizak/termui"
)

var (
	tools = []string{
		"Auslastung",
		"Deckungsbeitrag",
	}
	sskList *toolList
	reqList *requirementsList
	grid    *ui.Grid
	nav     *navigator
)

const ()

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	grid = ui.NewGrid()
	ui.Theme.Default = ui.NewStyle(ui.ColorBlack)

	sskList = newToolList(tools)
	reqList = newRequirementsList(map[string][]string{
		"Auslastung":      {"A 1", "A 2"},
		"Deckungsbeitrag": {"D 1", "D 2"},
	})

	nav = newNavigator(&[]navigatable{sskList, reqList}, grid)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0/2, sskList.list),
			ui.NewCol(1.0/2, reqList.list),
		),
	)

	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	ui.Render(grid)
	eventLoop()
}

func eventLoop() {

	// handles kill signal sent to gotop
	sigTerm := make(chan os.Signal, 2)
	signal.Notify(sigTerm, os.Interrupt, syscall.SIGTERM)

	uiEvents := ui.PollEvents()
	//previousKey := ""

	for {
		select {
		case <-sigTerm:
			return
		case e := <-uiEvents:
			switch e.ID {
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				termWidth, termHeight := payload.Width, payload.Height
				grid.SetRect(0, 0, termWidth, termHeight)
				ui.Clear()
				ui.Render(grid)
			case "q", "<C-c>":
				return
			case "<Down>":
				nav.down()

			case "<Up>":
				nav.up()

			case "<Tab>":
				nav.focusOnNextItem()
			}
		}
	}
}

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
	if nav.index > len(*nav.items) {
		nav.index = 0
		(*nav.items)[0].setActive(true)
		(*nav.items)[len(*nav.items)-1].setActive(false)
		ui.Render(nav.grid)
		return
	}
	(*nav.items)[nav.index+1].setActive(true)
	(*nav.items)[nav.index].setActive(false)
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

type toolList struct {
	list   *widgets.List
	tools  []string
	active bool
}

func newToolList(tools []string) *toolList {
	tl := toolList{
		list:   widgets.NewList(),
		tools:  tools,
		active: false,
	}
	tl.list.Rows = tools
	return &tl
}

func (tl *toolList) setTitle(title string) {
	tl.list.Title = title
}

func (tl *toolList) selectedRow() uint {
	return tl.list.SelectedRow
}

// satisfy navigatable
func (tl *toolList) up() {
	tl.list.ScrollUp()
}

func (tl *toolList) down() {
	tl.list.ScrollDown()
}

func (tl *toolList) setActive(active bool) {
	tl.active = active
	if active {
		tl.list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack)
	} else {
		tl.list.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)
	}
}

func (tl *toolList) selectedRowContent() string {
	return tl.list.Rows[tl.selectedRowIndex()]
}

func (tl *toolList) selectedRowIndex() uint {
	return tl.list.SelectedRow
}

type requirementsList struct {
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

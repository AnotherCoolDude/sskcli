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

	parser "github.com/AnotherCoolDude/sskcli/parser"
	wg "github.com/AnotherCoolDude/sskcli/widgets"

	ui "github.com/gizak/termui"
)

var (
	// tools = []string{
	// 	"Auslastung",
	// 	"Deckungsbeitrag",
	// }
	sskList  *wg.ToolList
	depList  *wg.DependencyList
	descList *wg.DependencyList
	grid     *ui.Grid
	nav      *wg.Navigator
)

const ()

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	initWidgets()
	setupGrid()
	eventLoop()
}

// setup / init
func setupGrid() {
	grid = ui.NewGrid()
	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, sskList.Griditem()),
			ui.NewCol(1.0/2, depList.Griditem()),
		),
		ui.NewRow(1.0/2, descList.Griditem()),
	)
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	ui.Render(grid)
}

func initWidgets() {
	config := parser.ParseYaml("./config/sskConfig.yaml")
	tools, deps, descs := config.Split()
	sskList = wg.NewToolList(tools)
	depList = wg.NewDependencyList(deps)
	// depList = wg.NewDependencyList(map[string][]string{
	// 	"Auslastung":      {"A 1", "A 2"},
	// 	"Deckungsbeitrag": {"D 1", "D 2"},
	// })
	descList = wg.NewDependencyList(descs)
	nav = wg.NewNavigator(&[]wg.Navigatable{sskList, depList, descList}, grid)
	depList.ListDependencies(sskList.SelectedRowContent())
	nav.FocusOnItem(0)
	nav.RenderItems()
}

// eventLoop
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
				nav.Down()
				if nav.FocusedItem() == sskList {
					depList.ListDependencies(sskList.SelectedRowContent())
					nav.RenderItems()
				}
			case "<Up>":
				nav.Up()
				if nav.FocusedItem() == sskList {
					depList.ListDependencies(sskList.SelectedRowContent())
					nav.RenderItems()
				}
			case "<Tab>":
				nav.FocusOnNextItem()
			}
		}
	}
}

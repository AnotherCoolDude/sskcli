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
	toolList = []string{
		"Auslastung",
		"Deckungsbeitrag",
	}
	sskList *widgets.List
)

const ()

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	grid := ui.NewGrid()
	sskList = initList()
	grid.Set(
		ui.NewRow(1.0/2, sskList),
		ui.NewRow(1.0/2, sskList),
	)

	ui.Render(grid)

	eventLoop()
}

func initList() *widgets.List {
	sskList := widgets.NewList()
	sskList.Title = "Selinka/Schmitz Toolbox"
	sskList.Rows = toolList
	sskList.TextStyle = ui.NewStyle(ui.ColorBlue)
	sskList.WrapText = false

	return sskList
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
			case "q", "<C-c>":
				return
			case "<Down>":
				sskList.ScrollDown()
			case "<Up>":
				sskList.ScrollUp()
			}
		}
	}
}

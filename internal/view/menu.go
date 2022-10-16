package view

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MenuView struct {
	*tview.Flex

	height int
}

type Matrix struct {
	*tview.Flex
	content *tview.TextView
}

func NewMenuView() *MenuView {
	view := &MenuView{
		Flex:   tview.NewFlex(),
		height: 10,
	}
	view.render()

	return view
}

func (m *MenuView) render() {
	m.Clear()

	matrix := &Matrix{
		Flex:    tview.NewFlex(),
		content: tview.NewTextView().SetDynamicColors(true),
	}
	matrix.SetBorder(true).
		SetBorderColor(tcell.ColorGreen).
		SetTitle("Menu").
		SetTitleColor(tcell.ColorGreen).
		SetTitleAlign(tview.AlignLeft)

	matrix.AddItem(matrix.content, 0, 1, true)
	m.fillup(matrix)

	width := 80
	m.AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 5, 1, false).
			AddItem(matrix, m.height+2, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func (m *MenuView) fillup(matrix *Matrix) {
	var text string
	switch {
	case choiceView.HasFocus():
		m.height = len(commandMenus)
		for _, menu := range commandMenus {
			text += fmt.Sprintf("%-6s%-s\n", menu.name, menu.desc)
		}
	case listView.HasFocus():
		m.height = len(listMenus)
		for _, menu := range listMenus {
			text += fmt.Sprintf("%-6s%-s\n", menu.name, menu.desc)
		}
	case searchView.HasFocus():
		m.height = len(searchMenus)
		for _, menu := range searchMenus {
			text += fmt.Sprintf("%-6s%-s\n", menu.name, menu.desc)
		}
	}

	matrix.content.SetText(text)
}

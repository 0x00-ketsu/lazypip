package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type WelcomeView struct {
	*tview.Flex
	blank *tview.TextView
	hint *tview.TextView
}

func NewWelcomeView() *WelcomeView {
	view := &WelcomeView{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
		blank: tview.NewTextView(),
		hint: tview.NewTextView().SetTextColor(tcell.ColorYellow).SetTextAlign(tview.AlignCenter),
	}
	view.SetBorder(true).SetTitle(" Lazypip ")
	view.AddItem(view.blank, 0, 1, false).
		AddItem(view.hint, 0, 1, false)

	view.hint.SetText("Welcome to Lazypip!\n---------------------\nA simple TUI interactive with Python pip command")

	return view
}

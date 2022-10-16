package view

import (
	"github.com/rivo/tview"
)

type ListDetailView struct {
	*tview.Flex
	content *tview.TextView
}

func NewListDetailView() *ListDetailView {
	view := &ListDetailView{
		Flex:    tview.NewFlex(),
		content: tview.NewTextView().SetDynamicColors(true),
	}
	view.AddItem(view.content, 0, 1, false)

	view.
		SetTitle(" Detail ").
		SetBorder(true).
		SetBorderPadding(1, 0, 1, 0)

	return view
}

func (l *ListDetailView) render(text string) {
	app.SetFocus(l)
	l.content.Clear().SetText(text)
}

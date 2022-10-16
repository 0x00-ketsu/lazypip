package view

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StatuslineView struct {
	*tview.Flex
	container *tview.Application
	hint      *tview.TextView
}

func NewStatuslineView(app *tview.Application) *StatuslineView {
	view := &StatuslineView{
		Flex: tview.NewFlex(),
		container: app,
		hint: tview.NewTextView().SetDynamicColors(true),
	}

	view.setDefaultHint()
	view.AddItem(view.hint, 0, 1, false)

	return view
}

func (s *StatuslineView) setDefaultHint() {
	s.hint.SetTextColor(tcell.ColorBlue).
	SetTextAlign(tview.AlignLeft).
		SetText("q: quit, x: menu, ↑/k ↓/j: navigate")
}

func (s *StatuslineView) restore() {
	s.container.QueueUpdateDraw(func() {
		s.setDefaultHint()
	})
}

// Used to skip queued restore of Status panel
// in case of new showForSeconds within waiting period
var restorInQ = 0

func (s *StatuslineView) showForSeconds(message string, duration int) {
	if s.container == nil {
		return
	}
	s.hint.SetText(message)
	restorInQ++

	go func() {
		time.Sleep(time.Second * time.Duration(duration))

		// Apply restore only if this is the last pending restore
		if restorInQ == 1 {
			s.restore()
		}

		restorInQ--
	}()
}

package view

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: Support command interaction

type CommandView struct {
	*tview.Flex
	input   *tview.InputField
	hint    *tview.TextView
	content *tview.TextView
}

func NewCommandView() *CommandView {
	view := &CommandView{
		Flex:    tview.NewFlex().SetDirection(tview.FlexRow),
		input:   tview.NewInputField(),
		hint:    tview.NewTextView().SetTextAlign(tview.AlignCenter).SetDynamicColors(true),
		content: tview.NewTextView(),
	}
	view.SetTitle(" Command ").
		SetBorder(true)

	view.input.
		SetLabel(" [::u]p[::-]ip ").SetLabelColor(tcell.ColorYellow).
		SetFieldWidth(60).
		SetPlaceholder("Input <command> [options]").
		SetBorderPadding(1, 0, 0, 0)
	view.content.SetBorder(true)

	view.AddItem(view.input, 2, 1, true).
		AddItem(view.hint, 1, 1, false)

	view.input.SetAutocompleteFunc(func(currentText string) (entries []string) {
		for _, command := range helpCommands {
			if strings.Contains(command, currentText) {
				entries = append(entries, command)
			}
		}
		return
	})
	view.input.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			text := view.input.GetText()
			view.execute(text)
			return
		case tcell.KeyEsc:
			app.SetFocus(aside)
		}
	})

	return view
}

func (c *CommandView) initial() {
	c.input.SetText("")
	c.clear()
}

func (c *CommandView) execute(text string) {
	if len(text) == 0 {
		c.setHint(" Install package should not be empty! ", tcell.ColorRed)
		return
	}

	c.openContent()

	args := strings.Split(text, " ")
	executePipCmdAync(args, commandView.content, nil)
}

func (c *CommandView) renderContent(text string, color tcell.Color) {
	c.content.SetText(text).SetTextColor(color)
}

func (c *CommandView) openContent() {
	c.clear()

	c.AddItem(c.content, 0, 1, false)
	c.content.SetChangedFunc(func() {
		app.Draw()
	})
}

func (c *CommandView) clear() {
	c.hint.Clear()
	c.content.Clear()

	c.RemoveItem(c.content)
}

func (c *CommandView) setHint(text string, color tcell.Color) {
	c.hint.SetText(text).SetTextColor(color)
}

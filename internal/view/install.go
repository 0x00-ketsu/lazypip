package view

import (
	"fmt"

	"github.com/0x00-ketsu/lazypip/internal/pip"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type InstallView struct {
	*tview.Flex
	input   *tview.InputField
	hint    *tview.TextView
	content *tview.TextView
}

func NewInstallView() *InstallView {
	view := &InstallView{
		Flex:    tview.NewFlex().SetDirection(tview.FlexRow),
		input:   tview.NewInputField(),
		hint:    tview.NewTextView().SetTextAlign(tview.AlignCenter).SetDynamicColors(true),
		content: tview.NewTextView(),
	}
	view.SetTitle(" Install ").
		SetBorder(true)

	view.input.SetFieldWidth(40).SetBorderPadding(1, 0, 0, 0)
	view.input.SetLabel(" [::u]i[::-]nstall: ").SetPlaceholder("Input Python package name ...")
	view.content.SetBorder(true)

	view.AddItem(view.input, 2, 1, true).
		AddItem(view.hint, 1, 1, false)

	view.input.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			name := view.input.GetText()
			view.install(name)
			return
		case tcell.KeyEsc:
			app.SetFocus(aside)
		}
	})

	return view
}

func (i *InstallView) initial() {
	i.input.SetText("")
	i.clear()
}

func (i *InstallView) install(name string) {
	if len(name) == 0 {
		i.setHint(" Install package should not be empty! ", tcell.ColorRed)
		return
	}

	i.openContent()

	installed, result := pip.IsPackageInstalled(name)
	if installed {
		i.setHint(fmt.Sprintf("Package: %s is already installed!", name), tcell.ColorRed)
		i.renderContent(result, tcell.ColorRed)
		return
	} else {
		i.content.SetTextColor(tcell.ColorWhite)
	}

	args := []string{"install", "--no-input", name}
	executePipCmdAync(args, installView.content, nil)
}

func (i *InstallView) renderContent(text string, color tcell.Color) {
	i.content.SetText(text).SetTextColor(color)
}

func (i *InstallView) openContent() {
	i.clear()

	i.AddItem(i.content, 0, 1, false)
	i.content.SetChangedFunc(func() {
		app.Draw()
	})
}

func (i *InstallView) clear() {
	i.hint.Clear()
	i.content.Clear()

	i.RemoveItem(i.content)
}

func (i *InstallView) setHint(text string, color tcell.Color) {
	i.hint.SetText(text).SetTextColor(color)
}

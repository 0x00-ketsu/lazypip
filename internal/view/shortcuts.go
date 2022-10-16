package view

import (
	"reflect"

	"github.com/0x00-ketsu/lazypip/internal/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Store focus view before open Menu view
var focusedView tview.Primitive

func setKeyboardShortcuts(app *tview.Application) *tview.Application {
	return app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if ignoreKeyEvent(app) {
			return event
		}

		// global shortcuts
		switch event.Rune() {
		case 'q':
			app.Stop()
			return nil
		case 'x':
			focusedView = app.GetFocus()
			menuView.render()
			pages.AddPage("menu", menuView, true, true)
			return nil
		}

		// views shortcuts
		switch {
		case listView.table.HasFocus():
			event = handleListResult(app, event)

		case menuView.HasFocus():
			event = handleMenuView(app, event)
		case choiceView.HasFocus():
			event = handleChoiceView(app, event)
		case listView.HasFocus():
			event = handleListView(app, event)
		case listDetailView.HasFocus():
			event = handleListDetailView(app, event)
		case searchView.HasFocus():
			event = handleSearchView(app, event)
		case installView.HasFocus():
			event = handleInstallView(app, event)
		case commandView.HasFocus():
			event = handleCommandView(app, event)
		}

		return event
	})
}

// Shortcuts for ListView table result
func handleListResult(app *tview.Application, event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		listView.blurTable()
		app.SetFocus(listView)
	case tcell.KeyRune:
		switch event.Rune() {
		case 'f':
			listView.blurTable()
			app.SetFocus(listView.filter)
			return nil
		case 'c':
			listView.checkPkg()
			return nil
		case 's':
			listView.showPkgDetail()
			return nil
		case 'u':
			listView.upgradePkg()
			return nil
		case 'U':
			listView.uninstallPkg()
			return nil
		}
	}

	return event
}

func handleMenuView(app *tview.Application, event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		pages.RemovePage("menu")
		app.SetFocus(focusedView)
		return nil
	}

	return event
}

func handleSearchView(app *tview.Application, event *tcell.EventKey) *tcell.EventKey {
	this := searchView

	switch event.Key() {
	case tcell.KeyEsc:
		if !this.HasFocus() {
			app.SetFocus(this)
		} else {
			app.SetFocus(aside)
		}
		return nil
	case tcell.KeyRune:
		switch event.Rune() {
		case 's':
			this.focusSearch()
			return nil
		case 't':
			this.focusTable()
			return nil
		case 'p':
			this.prevPage()
			return nil
		case 'n':
			this.nextPage()
			return nil
		}
	}

	return event
}

func handleChoiceView(app *tview.Application, event *tcell.EventKey) *tcell.EventKey {
	this := choiceView

	switch event.Key() {
	case tcell.KeyRune:
		switch event.Rune() {
		case 'j':
			this.lineDown()
			return nil
		case 'k':
			this.lineUp()
			return nil
		case 'g':
			this.list.SetCurrentItem(0)
			return nil
		case 'G':
			this.list.SetCurrentItem(this.list.GetItemCount() - 1)
			return nil
		}
	}
	return event
}

func handleListView(app *tview.Application, event *tcell.EventKey) *tcell.EventKey {
	this := listView

	switch event.Key() {
	case tcell.KeyEsc:
		app.SetFocus(aside)
		return nil
	case tcell.KeyRune:
		switch event.Rune() {
		case 'f':
			this.focusFilter()
			return nil
		case 't':
			this.focusTable()
			return nil
		case 'p':
			this.prevPage()
			return nil
		case 'n':
			this.nextPage()
			return nil
		case 'g':
			this.gotoFirstPage()
			return nil
		case 'G':
			this.gotoLastPage()
			return nil
		case 'r':
			this.reloadPackages()
			return nil
		}
	}

	return event
}

func handleListDetailView(app *tview.Application, event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		listView.removeDetail()
		app.SetFocus(listView.table)
		listView.focusTable()
		return nil
	}

	return event
}

func handleInstallView(app *tview.Application, event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyESC:
		app.SetFocus(aside)
		return nil
	}

	return event
}

func handleCommandView(app *tview.Application, event *tcell.EventKey)*tcell.EventKey  {
	this := commandView

	switch event.Key() {
	case tcell.KeyEsc:
		app.SetFocus(aside)
		return nil
	case tcell.KeyRune:
		switch event.Rune() {
		case 'p':
			app.SetFocus(this.input)
			return nil
		}
	}

	return event
}

func ignoreKeyEvent(app *tview.Application) bool {
	ignores := []string{"*tview.InputField"}

	return utils.InArray(reflect.TypeOf(app.GetFocus()).String(), ignores)
}

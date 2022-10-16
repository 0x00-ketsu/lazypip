package view

import (
	"fmt"

	"github.com/0x00-ketsu/lazypip/internal/pip"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var sheaders = []string{"Name", "Version", "Released", "Description"}

type SearchView struct {
	*tview.Flex
	input     *tview.InputField
	paginator *tview.Flex
	table     *tview.Table
	hint      *tview.TextView

	query string
	page  int
}

func NewSearchView() *SearchView {
	view := &SearchView{
		Flex:      tview.NewFlex().SetDirection(tview.FlexRow),
		input:     tview.NewInputField(),
		hint:      tview.NewTextView(),
		paginator: tview.NewFlex(),
		table:     tview.NewTable().SetBorders(true),

		page: 1,
	}
	view.SetTitle(" Search ").
		SetBorder(true)

	view.input.SetFieldWidth(40).SetBorderPadding(1, 0, 1, 0)
	view.input.SetLabel("[::u]s[::-]earch: ").SetPlaceholder("Input Python package name ...")
	view.hint.SetTextAlign(tview.AlignCenter).SetDynamicColors(true).SetBorderPadding(1, 0, 0, 0)

	view.AddItem(view.input, 2, 1, true).
		AddItem(view.hint, 2, 1, false).
		AddItem(view.paginator, 1, 1, false).
		AddItem(view.table, 0, 1, false)

	view.setPaginator()

	// Search input handler
	view.input.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			view.query = view.input.GetText()
			view.search(view.query, view.page)
			view.focusTable()
			return
		case tcell.KeyEsc:
			view.focusTable()
		}
	})

	return view
}

func (s *SearchView) init() {
	s.input.SetText("")
	s.paginator.Clear()
	s.table.Clear()
}

func (s *SearchView) search(query string, page int) {
	s.table.Clear()

	if len(query) == 0 {
		s.setHint(" Query content should not be empty! ", tcell.ColorRed)
		return
	}

	packages, reqURL, err := pip.Search(query, page)
	if err != nil {
		packages = []pip.Package{}
	}

	if len(packages) == 0 {
		s.setHint(" No Match Package(s) Found! ", tcell.ColorGray)
	} else {
		s.setHint("[::i]🐍 "+reqURL+" 🐍", tcell.ColorRed)
		s.setPaginator()
		s.setTableHeader()
		s.fillTableCells(packages)
	}
}

func (s *SearchView) setPaginator() {
	if len(s.query) == 0 {
		s.paginator.Clear().
			AddItem(blankCell, 0, 1, false)
	} else {
		currentPage := tview.NewTextView().SetText(fmt.Sprintf("current page: %d", s.page)).SetTextColor(tcell.ColorYellow)
		s.paginator.Clear().
			AddItem(blankCell, 1, 1, false).
			AddItem(makeButton(" [::u]p[::-]rev", tcell.ColorWhite, s.prevPage), 6, 1, false).
			AddItem(blankCell, 1, 1, false).
			AddItem(makeButton(" [::u]n[::-]ext", tcell.ColorWhite, s.nextPage), 6, 1, false).
			AddItem(blankCell, 1, 1, false).
			AddItem(currentPage, 0, 1, false)
	}
}

func (s *SearchView) prevPage() {
	oldPage := s.page

	if s.page > 1 {
		s.page--
	}

	if s.page != oldPage {
		s.search(s.query, s.page)
	}
}

func (s *SearchView) nextPage() {
	oldPage := s.page

	s.page++
	if s.page != oldPage {
		s.search(s.query, s.page)
	}
}

// Set search table header
func (s *SearchView) setTableHeader() {
	for idx, column := range sheaders {
		cell := tview.NewTableCell(" " + column + " ").SetAlign(tview.AlignLeft).SetSelectable(false)
		s.table.SetCell(0, idx, cell)
	}
}

// Fillup search table with Packages
func (s *SearchView) fillTableCells(packages []pip.Package) {
	for idx, pkg := range packages {
		name := tview.NewTableCell(PREFIX_PACKAGE + pkg.Name + ONE_SPACE).SetTextColor(tcell.ColorLightCyan).SetAlign(tview.AlignLeft)
		version := tview.NewTableCell(ONE_SPACE + pkg.Version + ONE_SPACE).SetTextColor(tcell.ColorOrange).SetAlign(tview.AlignLeft)
		released := tview.NewTableCell(ONE_SPACE + pkg.Released + ONE_SPACE).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft)
		desc := tview.NewTableCell(ONE_SPACE + pkg.Description + ONE_SPACE).SetTextColor(tcell.ColorPurple).SetAlign(tview.AlignLeft)

		s.table.SetCell(idx+1, 0, name)
		s.table.SetCell(idx+1, 1, version)
		s.table.SetCell(idx+1, 2, released)
		s.table.SetCell(idx+1, 3, desc)
	}
}

func (s *SearchView) blur() {
	s.table.SetSelectable(false, false)
	s.table.Blur()
}

func (s *SearchView) focusSearch() {
	app.SetFocus(s.input)

	if s.table.HasFocus() {
		s.table.SetSelectable(false, false)
		s.table.Blur()
	}
}

func (s *SearchView) focusTable() {
	app.SetFocus(s.table)
	s.table.SetSelectable(true, false)

	if s.input.HasFocus() {
		s.input.Blur()
	}
}

func (s *SearchView) setHint(text string, color tcell.Color) {
	s.hint.SetText(text).SetTextColor(color)
}

package view

import (
	"fmt"
	"math"
	"os/exec"
	"strings"

	"github.com/0x00-ketsu/lazypip/internal/pip"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var lheaders = []string{"Name", "Version"}

type ListView struct {
	*tview.Flex
	filter    *tview.InputField
	paginator *tview.Flex
	table     *tview.Table
	hint      *tview.TextView

	packages         []pip.Package
	filteredPackages []pip.Package

	page      int
	pageSize  int
	pageCount int // page total count
}

func NewListView() *ListView {
	view := &ListView{
		Flex:      tview.NewFlex().SetDirection(tview.FlexRow),
		filter:    tview.NewInputField().SetLabel("[::u]f[::-]ilter: ").SetPlaceholder("Input filter Python package name ..."),
		paginator: tview.NewFlex(),
		table:     tview.NewTable().SetBorders(true),
		hint:      tview.NewTextView(),

		page:     1,
		pageSize: conf.Pip.ListPageSize,
	}
	view.SetTitle(" List ").
		SetBorder(true)

	view.filter.SetFieldWidth(40).SetBorderPadding(1, 0, 1, 0)
	view.hint.SetTextAlign(tview.AlignCenter).SetDynamicColors(true)
	view.table.SetBorderPadding(0, 0, 1, 0)

	view.AddItem(view.filter, 2, 1, false).
		AddItem(view.hint, 1, 1, false).
		AddItem(view.paginator, 1, 1, false).
		AddItem(view.table, 0, 1, false)

	// Filter packages
	view.filter.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			filter := view.filter.GetText()
			view.filterPackage(filter)
			view.focusTable()

			app.SetFocus(view.table)
		case tcell.KeyEsc:
			view.table.SetSelectable(true, false)
			app.SetFocus(view.table)
		}
	})

	return view
}

func (l *ListView) initial() {
	l.clearHint()
	l.filter.SetText("")

	l.packages = pip.List()
	l.filteredPackages = l.packages

	l.loadPcakges()
	l.focusTable()
}

func (l *ListView) loadPcakges() {
	l.table.Clear()

	l.setTableHeader()
	l.setPaginator()
	l.fillTableCells()
}

func (l *ListView) reloadPackages() {
	l.packages = pip.List()
	l.filteredPackages = l.packages

	l.loadPcakges()
}

func (l *ListView) setPaginator() {
	pkgCount := len(l.filteredPackages)
	l.pageCount = int(math.Ceil(float64(pkgCount) / float64(l.pageSize)))

	currentPage := tview.NewTextView().
		SetText(fmt.Sprintf("Total: %d, page: %d/%d", pkgCount, l.page, l.pageCount)).
		SetTextColor(tcell.ColorYellow)
	l.paginator.Clear().
		AddItem(blankCell, 2, 1, false).
		AddItem(makeButton(" [::u]p[::-]rev", tcell.ColorWhite, l.prevPage), 6, 1, false).
		AddItem(blankCell, 1, 1, false).
		AddItem(makeButton(" [::u]n[::-]ext", tcell.ColorWhite, l.nextPage), 6, 1, false).
		AddItem(blankCell, 1, 1, false).
		AddItem(currentPage, 0, 1, false)
}

func (l *ListView) setTableHeader() {
	for idx, column := range lheaders {
		cell := tview.NewTableCell(" " + column + " ").SetAlign(tview.AlignLeft).SetSelectable(false)
		l.table.SetCell(0, idx, cell)
	}
}

// Fillup search table with Packages
func (l *ListView) fillTableCells() {
	for idx, pkg := range l.getPackages() {
		name := tview.NewTableCell(PREFIX_PACKAGE + pkg.Name + ONE_SPACE).SetTextColor(tcell.ColorLightCyan).SetAlign(tview.AlignLeft)
		version := tview.NewTableCell(ONE_SPACE + pkg.Version + ONE_SPACE).SetTextColor(tcell.ColorOrange).SetAlign(tview.AlignLeft)

		l.table.SetCell(idx+1, 0, name)
		l.table.SetCell(idx+1, 1, version)
	}
}

// Filter packages with case-insensitive
func (l *ListView) filterPackage(text string) {
	l.table.Clear()
	l.filteredPackages = nil

	if len(text) == 0 {
		l.filteredPackages = l.packages
	}

	for _, p := range l.packages {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(text)) {
			l.filteredPackages = append(l.filteredPackages, p)
		}
	}

	if len(l.filteredPackages) == 0 {
		l.paginator.Clear()
		l.setHint(" No Match Package(s) Found! ", tcell.ColorGray)
	} else {
		l.clearHint()
		l.loadPcakges()
	}
}

func (l *ListView) showPkgDetail() {
	l.openDetailView()

	var text string
	currentSelectPackage := l.getCurrentSelectPackage()
	out, err := exec.Command("pip", "show", currentSelectPackage).Output()
	if err != nil {
		text = ""
	}
	text = string(out[:])
	listDetailView.render(text)
}

func (l *ListView) checkPkg() {
	l.openDetailView()

	var text string
	currentSelectPackage := l.getCurrentSelectPackage()
	out, err := exec.Command("pip", "check", currentSelectPackage).Output()
	if err != nil {
		text = ""
	}
	text = string(out[:])
	listDetailView.render(text)
}

func (l *ListView) upgradePkg() {
	l.openDetailView()

	currentSelectPackage := l.getCurrentSelectPackage()
	args := []string{"install", "-U", currentSelectPackage}
	executePipCmdAync(args, listDetailView.content, l.reloadPackages)
}

// TODO: Add confirm modal
func (l *ListView) uninstallPkg() {
	l.openDetailView()

	currentSelectPackage := l.getCurrentSelectPackage()
	args := []string{"uninstall", "-y", currentSelectPackage}
	executePipCmdAync(args, listDetailView.content, l.reloadPackages)
}

func (l *ListView) openDetailView() {
	main.AddItem(listDetailView, 0, 2, false)

	listDetailView.content.SetChangedFunc(func() {
		app.Draw()
	})

	app.SetFocus(listDetailView)
	l.blurTable()
}

// Hide detail view
func (l *ListView) removeDetail() {
	main.RemoveItem(listDetailView)
}

func (l *ListView) blurTable() {
	l.table.SetSelectable(false, false)
	l.table.Blur()
}

func (l *ListView) focusFilter() {
	app.SetFocus(l.filter)

	if l.table.HasFocus() {
		l.table.SetSelectable(false, false)
		l.table.Blur()
	}
}

func (l *ListView) focusTable() {
	app.SetFocus(l.table)
	l.table.Select(1, 0)
	l.table.SetSelectable(true, false)

	if l.filter.HasFocus() {
		l.filter.Blur()
	}
}

// Return current select package in table
func (l *ListView) getCurrentSelectPackage() string {
	cellText := l.table.GetCell(l.table.GetSelection()).Text
	cellText = strings.TrimLeft(cellText, PREFIX_PACKAGE)
	return strings.TrimRight(cellText, ONE_SPACE)
}

func (l *ListView) setHint(text string, color tcell.Color) {
	l.hint.SetText(text).SetTextColor(color)
}

func (l *ListView) clearHint() {
	l.hint.Clear()
}

func (l *ListView) prevPage() {
	oldPage := l.page

	if l.page > 1 {
		l.page--
	}

	if l.page != oldPage {
		l.loadPcakges()
	}
}

func (l *ListView) nextPage() {
	oldPage := l.page

	l.page++
	if l.page != oldPage && l.page <= l.pageCount {
		l.loadPcakges()
	}
}

func (l *ListView) gotoFirstPage() {
	l.page = 1
	l.loadPcakges()
}

func (l *ListView) gotoLastPage() {
	l.page = l.pageCount
	l.loadPcakges()
}

func (l *ListView) getPackages() []pip.Package {
	var end int
	start := (l.page - 1) * l.pageSize
	if l.page == l.pageCount {
		end = start + (len(l.filteredPackages) - start)
	} else {
		end = l.page * l.pageSize
	}

	return l.filteredPackages[start:end]
}

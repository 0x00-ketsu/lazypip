package view

import (
	"github.com/0x00-ketsu/lazypip/internal/config"
	"github.com/0x00-ketsu/lazypip/internal/contrib/logging"
	"github.com/0x00-ketsu/lazypip/internal/pip"
	"github.com/rivo/tview"
	"go.uber.org/zap"
)

var (
	conf   *config.Config
	logger *zap.Logger
)

var (
	app                 *tview.Application
	layout, main, aside *tview.Flex
	pages               *tview.Pages

	choiceView     *ChoiceView
	listView       *ListView
	listDetailView *ListDetailView
	searchView     *SearchView
	installView    *InstallView
	commandView    *CommandView
	menuView       *MenuView
	statuslineView *StatuslineView

	helpCommands []string
)

const (
	PREFIX_PACKAGE = "📂 "
	ONE_SPACE      = " "
)

func init() {
	conf, _ = config.Load()
	logger = logging.GetLogger()
}

func Load(application *tview.Application) *tview.Flex {
	app = application

	helpCommands = pip.GetCommands()

	choiceView = NewChoiceView()
	listView = NewListView()
	listDetailView = NewListDetailView()
	searchView = NewSearchView()
	installView = NewInstallView()
	commandView = NewCommandView()
	statuslineView = NewStatuslineView(app)

	// aside window
	aside = tview.NewFlex().SetDirection(tview.FlexRow)
	aside.AddItem(choiceView, 0, 1, true)

	// main window
	main = tview.NewFlex()
	main.AddItem(aside, 35, 1, true).
		AddItem(NewWelcomeView(), 0, 1, false)

	// use pages show menu modal
	menuView = NewMenuView()
	pages = tview.NewPages().
		AddPage("main", main, true, true).
		AddPage("menu", menuView, true, false)

	// layout
	layout = tview.NewFlex().SetDirection(tview.FlexRow)
	layout.AddItem(pages, 0, 1, true).
		AddItem(statuslineView, 1, 1, false)

	// set keyboard shortcuts
	setKeyboardShortcuts(app)

	return layout
}

// Load search view in main window
func loadSearchView() {
	onlyAsideInMain()

	main.AddItem(searchView, 0, 1, false)
	app.SetFocus(searchView)

	searchView.init()
}

func loadListView() {
	onlyAsideInMain()

	main.AddItem(listView, 0, 1, false)
	app.SetFocus(listView)

	listView.initial()
}

func loadInstallView() {
	onlyAsideInMain()

	main.AddItem(installView, 0, 1, true)
	app.SetFocus(installView)

	installView.initial()
}

func loadCommandView() {
	onlyAsideInMain()

	main.AddItem(commandView, 0, 1, true)
	app.SetFocus(commandView)

	commandView.initial()
}

// Keep only aside window in main window
// Remove all other views from main window
func onlyAsideInMain() {
	itemCnt := main.GetItemCount()
	if itemCnt > 1 {
		for i := 1; i < itemCnt; i++ {
			item := main.GetItem(i)
			main.RemoveItem(item)
		}
	}
}

package view

import "github.com/rivo/tview"

type Choice struct {
	text     string
	desc     string
	shortcut rune
	handler  func()
}

var choices = []Choice{
	{
		text:     "search",
		desc:     "Search PyPI for packages",
		shortcut: 's',
		handler:  loadSearchView,
	},
	{
		text:     "list",
		desc:     "List installed packages",
		shortcut: 'l',
		handler:  loadListView,
	},
	{
		text:     "install",
		desc:     "Install packages",
		shortcut: 'i',
		handler:  loadInstallView,
	},
	{
		text:     "command",
		desc:     "Execute pure pip commands",
		shortcut: 'c',
		handler:  loadCommandView,
	},
}

type ChoiceView struct {
	*tview.Flex
	list *tview.List
}

func NewChoiceView() *ChoiceView {
	view := &ChoiceView{
		Flex: tview.NewFlex(),
		list: tview.NewList(),
	}
	view.AddItem(view.list, 0, 1, true)
	view.SetTitle(" Choices ").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	for _, choice := range choices {
		view.list.AddItem(choice.text, choice.desc, choice.shortcut, choice.handler)
	}

	return view
}

// Move to next list item
func (c *ChoiceView) lineDown() {
	curItemIdx := c.list.GetCurrentItem()
	itemCnt := c.list.GetItemCount()
	if curItemIdx >= 0 && curItemIdx < itemCnt-1 {
		nextItemIdx := curItemIdx + 1
		c.list.SetCurrentItem(nextItemIdx)
	}
}

// Move to previous list item
func (c *ChoiceView) lineUp() {
	curItemIdx := c.list.GetCurrentItem()
	itemCnt := c.list.GetItemCount()
	if curItemIdx < itemCnt && curItemIdx > 0 {
		prevItemIdx := curItemIdx - 1
		c.list.SetCurrentItem(prevItemIdx)
	}
}

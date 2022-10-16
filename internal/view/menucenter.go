package view

// For views register self menu
type menuItem struct {
	name string
	desc string
}

// Command view menus
var commandMenus = []menuItem{
	{name: "g", desc: "go to first item"},
	{name: "G", desc: "go to last item"},
}

// List view menus
var listMenus = []menuItem{
	{name: "f", desc: "filter packages"},
	{name: "t", desc: "focus on table"},
	{name: "p", desc: "previous page"},
	{name: "n", desc: "next page"},
	{name: "g", desc: "go to first page"},
	{name: "G", desc: "go to last page"},
	{name: "", desc: ""},
	{name: "r", desc: "refresh packages"},
	{name: "c", desc: "check selected package have compatible dependencies (in result table)"},
	{name: "s", desc: "show information about selected package (in result table)"},
	{name: "u", desc: "upgrade selected package (in result table)"},
	{name: "U", desc: "uninstall selected package (in result table)"},
}

// Search view menus
var searchMenus = []menuItem{
	{name: "s", desc: "search packages"},
	{name: "t", desc: "focus on table"},
	{name: "p", desc: "previous page"},
	{name: "n", desc: "next page"},
}

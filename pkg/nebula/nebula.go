package nebula

import (
	"embed"
	"net/http"

	"github.com/sonrhq/sonr/pkg/nebula/ui"
)

// TODO: Add CDN package for Shoelace as go-embed
//	labels: HTMX/Frontend,Plane,Github
//	milestone: 24

//go:embed assets/*
var assets embed.FS

// ServeAssets serves the assets from the embed.FS including stylesheets, images, and javascript files.
func ServeAssets() (pattern string, handler http.Handler) {
	return "/*", http.FileServer(http.FS(assets))
}

// TODO: Create Props and Slots interfaces
//	labels: HTMX/Frontend,Plane,Github
//	milestone: 24

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Aliases to UI Components                            ||
// ! ||--------------------------------------------------------------------------------||

// TODO: Implement Remaining UI Components from Shoelace
//	labels: HTMX/Frontend,Plane,Github
//	milestone: 24

// Accordian is an alias to the UI component.
var Accordian = ui.Accordian

// AlertDialog is an alias to the UI component.
var AlertDialog = ui.AlertDialog

// Alert is an alias to the UI component.
var Alert = ui.Alert

// Avatar is an alias to the UI component.
var Avatar = ui.Avatar

// Badge is an alias to the UI component.
var Badge = ui.Badge

// Breadcrumb is an alias to the UI component.
var Breadcrumb = ui.Breadcrumb

// Button is an alias to the UI component.
var Button = ui.Button

// ButtonGroup is an alias to the UI component.
var ButtonGroup = ui.ButtonGroup

// Card is an alias to the UI component.
var Card = ui.Card

// Carousel is an alias to the UI component.
var Carousel = ui.Carousel

// ChartArea is an alias to the UI component.
var ChartArea = ui.ChartArea

// ChartBar is an alias to the UI component.
var ChartBar = ui.ChartBar

// ChartLine is an alias to the UI component.
var ChartLine = ui.ChartLine

// ChartPie is an alias to the UI component.
var ChartPie = ui.ChartPie

// Checkbox is an alias to the UI component.
var Checkbox = ui.Checkbox

// Combobox is an alias to the UI component.
var Combobox = ui.Combobox

// CommandMenu is an alias to the UI component.
var CommandMenu = ui.CommandMenu

// ContextMenu is an alias to the UI component.
var ContextMenu = ui.ContextMenu

// DataTable is an alias to the UI component.
var DataTable = ui.DataTable

// DatePicker is an alias to the UI component.
var DatePicker = ui.DatePicker

// Dialog is an alias to the UI component.
var Dialog = ui.Dialog

// Divider is an alias to the UI component.
var Divider = ui.Divider

// Drawer is an alias to the UI component.
var Drawer = ui.Drawer

// Dropdown is an alias to the UI component.
var Dropdown = ui.Dropdown

// HoverCard is an alias to the UI component.
var HoverCard = ui.HoverCard

// Input is an alias to the UI component.
var Input = ui.Input

// Menubar is an alias to the UI component.
var Menubar = ui.Menubar

// NavigationMenu is an alias to the UI component.
var NavigationMenu = ui.NavigationMenu

// Pagination is an alias to the UI component.
var Pagination = ui.Pagination

// Popover is an alias to the UI component.
var Popover = ui.Popover

// Progress is an alias to the UI component.
var Progress = ui.Progress

// RadioGroup is an alias to the UI component.
var RadioGroup = ui.RadioGroup

// Resizeable is an alias to the UI component.
var Resizeable = ui.Resizeable

// ScrollArea is an alias to the UI component.
var ScrollArea = ui.ScrollArea

// Select is an alias to the UI component.
var Select = ui.Select

// Sheet is an alias to the UI component.
var Sheet = ui.Sheet

// Skeleton is an alias to the UI component.
var Skeleton = ui.Skeleton

// Slider is an alias to the UI component.
var Slider = ui.Slider

// Sonner is an alias to the UI component.
var Sonner = ui.Sonner

// Switch is an alias to the UI component.
var Switch = ui.Switch

// Tabs is an alias to the UI component.
var Tabs = ui.Tabs

// TabPanel is an alias to the UI component.
var TabPanel = ui.TabPanel

// Table is an alias to the UI component.
var Table = ui.Table

// TextArea is an alias to the UI component.
var TextArea = ui.TextArea

// Toast is an alias to the UI component.
var Toast = ui.Toast

// ToggleGroup is an alias to the UI component.
var ToggleGroup = ui.ToggleGroup

// Tooltip is an alias to the UI component.
var Tooltip = ui.Tooltip

// TextH1 is an alias to the UI component.
var TextH1 = ui.TextH1

// TextH2 is an alias to the UI component.
var TextH2 = ui.TextH2

// TextH3 is an alias to the UI component.
var TextH3 = ui.TextH3

// TextBody is an alias to the UI component.
var TextBody = ui.TextBody

// TextLink is an alias to the UI component.
var TextLink = ui.TextLink

// TextSmall is an alias to the UI component.
var TextSmall = ui.TextSmall

// TextCaption is an alias to the UI component.
var TextCaption = ui.TextCaption

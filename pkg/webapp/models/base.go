package models

type NavHeader struct {
	Logo    *Image
	Primary *NavItem
	Items   []*NavItem
}

type NavItem struct {
	Text string
	Href string
}

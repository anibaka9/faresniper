package main

type SidebarItem struct {
	Title    string
	Link     string
	IsActive bool
}

var sidebar = []SidebarItem{
	{
		Title: "Dashboard",
		Link:  "/",
	},
	{
		Title: "Airlines",
		Link:  "/airlines",
	},
	{
		Title: "Airports",
		Link:  "/airports",
	},
	{
		Title: "Cities",
		Link:  "/cities",
	},
	{
		Title: "Countries",
		Link:  "/countries",
	},
}

func GetSidebarData(activePage string) []SidebarItem {
	for i := range sidebar {
		if activePage == sidebar[i].Link {
			sidebar[i].IsActive = true
		} else {
			sidebar[i].IsActive = false
		}
	}
	return sidebar
}

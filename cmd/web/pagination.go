package main

import (
	"strconv"
)

type PaginationParams struct {
	TotalPages   int64
	TotalCount   int64
	CurrentPage  int64
	CurrentLimit int64
}

type PaginationData struct {
	Pages    []PageItem
	HasPrev  bool
	HasNext  bool
	PrevPage int
	NextPage int
}

type PageItem struct {
	Label    string
	Number   int
	IsActive bool
	IsGap    bool
}

func GetPaginationData(pagination PaginationParams) PaginationData {
	total := pagination.TotalPages
	current := pagination.CurrentPage

	var pages []PageItem

	if total <= 5 {
		for i := int64(0); i < total; i++ {
			pages = append(pages, PageItem{
				Number:   int(i),
				Label:    strconv.Itoa(int(i) + 1),
				IsActive: i == current,
			})
		}
	} else {
		pages = append(pages, PageItem{Number: 0, Label: "1", IsActive: current == 0})
		pages = append(pages, PageItem{Number: 1, Label: "2", IsActive: current == 1})
		pages = append(pages, PageItem{Number: 0, Label: "...", IsGap: true})
		pages = append(pages, PageItem{Number: int(total - 2), Label: strconv.Itoa(int(total - 1)), IsActive: current == total-2})
		pages = append(pages, PageItem{Number: int(total - 1), Label: strconv.Itoa(int(total)), IsActive: current == total-1})
	}

	return PaginationData{
		Pages:    pages,
		HasPrev:  current != 0,
		HasNext:  current != total-1,
		PrevPage: int(current) - 1,
		NextPage: int(current) + 1,
	}
}

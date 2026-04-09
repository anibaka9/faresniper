package main

import (
	"fmt"
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

const MaxPages = 4

func GetPaginationData(pagination PaginationParams) PaginationData {
	fmt.Println(pagination.TotalPages)
	pagesNum := 5
	if pagination.TotalPages < 5 {
		pagesNum = int(pagination.TotalPages)
	}
	pages := make([]PageItem, pagesNum)
	if pagination.TotalPages <= MaxPages {
		for i := 0; i < int(pagination.TotalPages); i++ {
			pages[i] = PageItem{
				Number:   i,
				Label:    strconv.Itoa(i + 1),
				IsActive: i == int(pagination.CurrentPage),
				IsGap:    false,
			}
		}
	} else {
		for i := range MaxPages / 2 {
			pages[i] = PageItem{
				Number:   i,
				Label:    strconv.Itoa(i + 1),
				IsActive: i == int(pagination.CurrentPage),
				IsGap:    false,
			}
		}
		pages[MaxPages/2] = PageItem{
			Number:   MaxPages/2 - 1,
			Label:    strconv.Itoa(MaxPages/2 + 1),
			IsActive: false,
			IsGap:    true,
		}
		for i := int(pagination.TotalPages - MaxPages/2); i < int(pagination.TotalPages); i++ {
			pages[i] = PageItem{
				Number:   i,
				Label:    strconv.Itoa(i + 1),
				IsActive: i == int(pagination.CurrentPage),
				IsGap:    false,
			}
		}
	}
	return PaginationData{
		Pages:    pages,
		HasPrev:  pagination.CurrentPage != 0,
		HasNext:  pagination.CurrentPage != pagination.TotalPages-1,
		PrevPage: int(pagination.CurrentPage) - 1,
		NextPage: int(pagination.CurrentPage) + 1,
	}
}

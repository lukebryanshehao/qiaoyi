package model

type Page struct {
	PageSize int
	PageIndex int
	PageCount int
}

func PageIndex(page *Page) int {
	if page.PageIndex == 0 {
		page.PageIndex = 1
	}
	return (page.PageIndex - 1) * page.PageSize
}
package page

import "strconv"

type Page struct {
	page  int
	limit int
}

// With initializes Page with given limit at page 1
func With(limit int) Page {
	return Page{page: 1, limit: limit}
}

// At sets page number - pages start at 1
func (p Page) At(page int) Page {
	p.page = page
	return p
}

func (p Page) Offset() int {
	return (p.page - 1) * p.limit
}

func (p Page) Limit() int {
	return p.limit
}

func (p Page) IsValid() bool {
	return p.page >= 0 && p.limit > 0
}

type getter interface {
	Get(key string) string
}

func From(g getter) (Page, error) {

	limitStr := g.Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return Page{}, err
	}

	pageStr := g.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return Page{}, err
	}

	return Page{
		page:  page,
		limit: limit,
	}, nil
}

func FromWithDefault(g getter, defaultPage Page) Page {
	page, err := From(g)
	if err != nil {
		return defaultPage
	}
	return page
}

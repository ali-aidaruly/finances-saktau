package sort

import (
	"errors"
	"fmt"
	"strings"
)

type direction int

const (
	asc  direction = 1
	desc direction = -1
)

type Sort struct {
	field     string
	direction direction
}

// By initializes Sort with given field in ASC order
func By(field string) Sort {
	return Sort{field: field, direction: asc}
}

func (s Sort) ASC() Sort {
	s.direction = asc
	return s
}

func (s Sort) DESC() Sort {
	s.direction = desc
	return s
}

func (s Sort) Field() string {
	// FIXME: find a way to escape the field
	return s.field
}

func (s Sort) IsASC() bool {
	return s.direction == asc
}

func (s Sort) IsDESC() bool {
	return s.direction == desc
}

func (s Sort) SQLDir() string {
	if s.IsASC() {
		return "ASC"
	}
	return "DESC"
}

func (s Sort) IsValid() bool {
	return s.field != "" && (s.direction == asc || s.direction == desc)
}

type getter interface {
	Get(key string) string
}

var (
	ErrInvalidSortKey   = errors.New("invalid sort key")
	ErrInvalidSortOrder = errors.New("invalid sort order")
)

func From(g getter, validSortKeys map[string]struct{}) (Sort, error) {

	field := g.Get("sort")
	if field == "" {
		return Sort{}, nil
	}

	if _, ok := validSortKeys[field]; !ok {
		return Sort{}, fmt.Errorf("%w: %s", ErrInvalidSortKey, field)
	}

	sort := Sort{field: field}

	dir := g.Get("order")
	switch strings.ToLower(dir) {
	case "asc":
		sort.direction = asc
	case "desc":
		sort.direction = desc
	default:
		return Sort{}, fmt.Errorf("%w: %s", ErrInvalidSortOrder, dir)
	}

	return sort, nil
}

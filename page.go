// Package page helps to manage pagination.
// While many of these functions may be considered trivial, they are often
// implemented incorrectly.
//
// Inspired by Leon Brocard's Data::Page Perl module.
// http://search.cpan.org/perldoc?Data%3A%3APage
package page

import (
	"errors"
	"fmt"
	"math"
)

// P is a struct to help manage pagination.
// Instantiate it directly, e.g.
//    p := page.P{10, 50, 1}
type P struct {
	EntriesPerPage, TotalEntries, CurrentPage int
}

// Valid returns an error if:
//     - Any of EntriesPerPage, TotalEntries or CurrentPage is zero or negative.
//     - CurrentPage is greater than TotalEntries / EntriesPerPage.
func (p P) Valid() error {
	switch {
	case p.EntriesPerPage <= 0:
		return errors.New("invalid EntriesPerPage")
	case p.TotalEntries <= 0:
		return errors.New("invalid TotalEntries")
	case p.CurrentPage <= 0:
		return errors.New("invalid CurrentPage")
	}
	if p.CurrentPage > p.TotalEntries/p.EntriesPerPage {
		return errors.New("out of range: CurrentPage")
	}
	return nil
}

// EntriesOnThisPage returns the number of entries on the current page.
func (p P) EntriesOnThisPage() int {
	return p.Last() - p.First() + 1
}

// FirstPage returns the first page.
// This will always be 1, but is present for symmetry with LastPage.
func (p P) FirstPage() int {
	return 1
}

// LastPage returns the total number of pages.
func (p P) LastPage() int {
	pages, rem := math.Modf(float64(p.TotalEntries) / float64(p.EntriesPerPage))
	lastPage := int(pages)
	if rem != 0 {
		lastPage++
	}
	if lastPage < 1 {
		lastPage = 1
	}
	return lastPage
}

// First returns the index of the first entry on the current page.
func (p P) First() int {
	if p.TotalEntries == 0 {
		return 0
	}
	return ((p.CurrentPage - 1) * p.EntriesPerPage) + 1
}

// Last returns the index of the last entry on the current page.
func (p P) Last() int {
	if p.CurrentPage == p.LastPage() {
		return p.TotalEntries
	}
	return p.CurrentPage * p.EntriesPerPage
}

// PrevPage returns the page before the current one.
// If the current page is the first one, the current page will be returned.
func (p P) PrevPage() int {
	if p.CurrentPage == 1 {
		return 1
	}
	return p.CurrentPage - 1
}

// NextPage returns the pager after the current one.
// If the current page is the last one, the current page will be returned.
func (p P) NextPage() int {
	if p.CurrentPage < p.LastPage() {
		return p.CurrentPage + 1
	}
	return p.CurrentPage
}

// String implements fmt.Stringer.
func (p P) String() string {
	return fmt.Sprintf("[epp:%d tot:%d cur:%d]", p.EntriesPerPage, p.TotalEntries, p.CurrentPage)
}

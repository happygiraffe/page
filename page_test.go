package page

import (
	"testing"
)

func TestString(t *testing.T) {
	p := P{10, 100, 1}
	if got, want := p.String(), "[epp:10 tot:100 cur:1]"; got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		p       P
		wantErr bool
	}{
		{P{-1, -1, -1}, true},
		{P{1, -1, -1}, true},
		{P{1, 1, -1}, true},
		{P{0, 0, 0}, true},
		{P{1, 0, 0}, true},
		{P{1, 1, 0}, true},
		{P{1, 1, 1}, false},
		// max CurrentPage is 2.
		{P{EntriesPerPage: 2, TotalEntries: 4, CurrentPage: 3}, true},
	}
	for i, tc := range tests {
		if err := tc.p.Valid(); (err != nil) != tc.wantErr {
			t.Errorf("%d. Valid() returned [%v], want error? %v", i, err, tc.wantErr)
		}
	}
}

func TestLastPage(t *testing.T) {
	tests := []struct {
		p    P
		want int
	}{
		{P{1, 1, 1}, 1},
		{P{1, 2, 1}, 2},
		{P{2, 6, 1}, 3},
		{P{2, 8, 4}, 4},
		{P{2, 7, 5}, 4},
	}
	for i, tc := range tests {
		if got := tc.p.LastPage(); got != tc.want {
			t.Errorf("%d. LastPage() = %d want %d", i, got, tc.want)
		}
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		p    P
		want int
	}{
		{P{1, 1, 1}, 1},
		{P{1, 2, 1}, 1},
		{P{2, 6, 1}, 1},
		{P{2, 8, 4}, 7},
		{P{2, 7, 5}, 9},
	}
	for i, tc := range tests {
		if got := tc.p.First(); got != tc.want {
			t.Errorf("%d. First() = %d want %d", i, got, tc.want)
		}
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		p    P
		want int
	}{
		{P{1, 1, 1}, 1},
		{P{2, 2, 1}, 2},
		{P{3, 2, 1}, 2},
		{P{3, 3, 1}, 3},
		{P{3, 6, 2}, 6},
	}
	for i, tc := range tests {
		if got := tc.p.Last(); got != tc.want {
			t.Errorf("%d. Last() = %d, want %d", i, got, tc.want)
		}
	}
}

func TestPrevPage(t *testing.T) {
	tests := []struct {
		p    P
		want int
	}{
		{P{1, 1, 1}, 1},
		{P{2, 2, 1}, 1},
		{P{2, 4, 2}, 1},
		{P{2, 6, 3}, 2},
		{P{3, 12, 4}, 3},
	}
	for i, tc := range tests {
		if got := tc.p.PrevPage(); got != tc.want {
			t.Errorf("%d. PrevPage() = %d, want %d", i, got, tc.want)
		}
	}
}

func TestNextPage(t *testing.T) {
	tests := []struct {
		p    P
		want int
	}{
		{P{1, 1, 1}, 1},
		{P{1, 2, 1}, 2},
		{P{2, 2, 1}, 1},
		{P{2, 6, 1}, 2},
		{P{2, 6, 2}, 3},
		{P{2, 6, 3}, 3},
	}
	for i, tc := range tests {
		t.Logf("%d. p = %v", i, tc.p)
		if got := tc.p.NextPage(); got != tc.want {
			t.Errorf("%d. NextPage() = %d, want %d", i, got, tc.want)
		}
	}
}

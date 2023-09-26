package mock

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCondMatcher_Matches(t *testing.T) {
	m := Cond(func(val int) bool {
		return val == 3

	})
	require.False(t, m.Matches("2"))
}

func TestInAnyOrder(t *testing.T) {
	tests := []struct {
		name      string
		wanted    []int
		given     any
		wantMatch bool
	}{
		{
			name:      "match for equal slices",
			wanted:    []int{1, 2, 3},
			given:     []int{1, 2, 3},
			wantMatch: true,
		},
		{
			name:      "match for slices with same elements of different order",
			wanted:    []int{1, 2, 3},
			given:     []int{1, 3, 2},
			wantMatch: true,
		},
		{
			name:      "not match for slices with different elements",
			wanted:    []int{1, 2, 3},
			given:     []int{1, 2, 4},
			wantMatch: false,
		},
		{
			name:      "not match for slices with missing elements",
			wanted:    []int{1, 2, 3},
			given:     []int{1, 2},
			wantMatch: false,
		},
		{
			name:      "not match for slices with extra elements",
			wanted:    []int{1, 2, 3},
			given:     []int{1, 2, 3, 4},
			wantMatch: false,
		},
		{
			name:      "match for empty slices",
			wanted:    []int{},
			given:     []int{},
			wantMatch: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InAnyOrder(tt.wanted).Matches(tt.given); got != tt.wantMatch {
				t.Errorf("got = %v, wantMatch %v", got, tt.wantMatch)
			}
		})
	}
}

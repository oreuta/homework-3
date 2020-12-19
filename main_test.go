package main

import (
	"reflect"
	"testing"
)

func TestTransform(t *testing.T) {
	testcases := []struct {
		name     string
		matrix   [][]int
		expSlice []int
		expError error
	}{
		{
			name:     "not ok, zero size slice",
			matrix:   [][]int{},
			expSlice: nil,
			expError: errZeroLength,
		},
		{
			name:     "ok, very small slice",
			matrix:   [][]int{[]int{1}},
			expSlice: []int{1},
			expError: nil,
		},
		{
			name: "ok, small slice",
			matrix: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			expSlice: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
			expError: nil,
		},
		{
			name: "ok, bigger slice",
			matrix: [][]int{
				[]int{1, 2, 3, 1},
				[]int{4, 5, 6, 4},
				[]int{7, 8, 9, 7},
				[]int{7, 8, 9, 7},
			},
			expSlice: []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8},
			expError: nil,
		},
		{
			name: "very not okey, egyptian pyramid",
			matrix: [][]int{
				[]int{1},
				[]int{2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9, 10},
			},
			expSlice: nil,
			expError: errInvalidSize,
		},
		{
			name: "not okey, not all rows have equal length",
			matrix: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5},
				[]int{6, 7, 8},
			},
			expSlice: nil,
			expError: errInvalidSize,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			slice, err := transform(tc.matrix)

			if err != tc.expError {
				t.Errorf("Expected err was \"%v\", got \"%v\"", tc.expError, err)
			}
			if !reflect.DeepEqual(slice, tc.expSlice) {
				t.Errorf("Expected slice was %v, got %v", tc.expSlice, slice)
			}
		})
	}
}

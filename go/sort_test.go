package main

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	tests := []struct {
		name      string
		givenNums []int
		wantNums  []int
	}{
		{
			"Empty slice",
			[]int{},
			[]int{},
		}, {
			"One number",
			[]int{5},
			[]int{5},
		}, {
			"Standard set",
			[]int{7, 5, 2, 3, 4},
			[]int{2, 3, 4, 5, 7},
		}, {
			"Duplicates",
			[]int{5, 5, 7, 7, 1, 1, 3, 2, 5},
			[]int{1, 1, 2, 3, 5, 5, 5, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNums := mergeSort(tt.givenNums); !reflect.DeepEqual(gotNums, tt.wantNums) {
				t.Errorf("mergeSort() = %v, want %v", gotNums, tt.wantNums)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name       string
		args       args
		wantSorted []int
	}{
		{
			"One empty",
			args{
				left:  nil,
				right: []int{2, 4, 6, 8},
			},
			[]int{2, 4, 6, 8},
		}, {
			"Two even lists",
			args{
				[]int{0, 3, 5},
				[]int{1, 2, 4},
			},
			[]int{0, 1, 2, 3, 4, 5},
		}, {
			"Two uneven lists",
			args{
				left:  []int{0, 3, 4},
				right: []int{1, 2},
			},
			[]int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSorted := merge(tt.args.left, tt.args.right); !reflect.DeepEqual(gotSorted, tt.wantSorted) {
				t.Errorf("merge() = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}

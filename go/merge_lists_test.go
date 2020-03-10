package main

import (
	"reflect"
	"testing"
)

func Test_mergeCookieLists(t *testing.T) {
	type args struct {
		mine  []int
		alice []int
	}
	tests := []struct {
		name       string
		args       args
		wantSorted []int
	}{
		{
			"Even lengths",
			args{
				mine:  []int{3, 5, 8},
				alice: []int{1, 4, 6},
			},
			[]int{1, 3, 4, 5, 6, 8},
		}, {
			"Alice has longer list",
			args{
				mine:  []int{3},
				alice: []int{5, 8, 12},
			},
			[]int{3, 5, 8, 12},
		}, {
			"I have longer list",
			args{
				mine:  []int{2, 4, 7, 8, 10},
				alice: []int{1, 3, 5, 6},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 10},
		}, {
			"Duplicates",
			args{
				mine:  []int{1, 3, 5, 6},
				alice: []int{1, 5, 6, 8},
			},
			[]int{1, 1, 3, 5, 5, 6, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSorted := mergeCookieLists(tt.args.mine, tt.args.alice); !reflect.DeepEqual(gotSorted, tt.wantSorted) {
				t.Errorf("mergeCookieLists() = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}

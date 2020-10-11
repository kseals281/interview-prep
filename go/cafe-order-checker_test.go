package main

import "testing"

func Test_isFIFO(t *testing.T) {
	type args struct {
		takeOut []int
		dineIn  []int
		served  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"empty",
			args{
				takeOut: []int{},
				dineIn:  []int{},
				served:  []int{},
			},
			true,
		}, {
			"correct order - take out first",
			args{
				takeOut: []int{1, 2},
				dineIn:  []int{3, 4},
				served:  []int{1, 2, 3, 4},
			},
			true,
		}, {
			"correct order - dine in first",
			args{
				takeOut: []int{1, 2},
				dineIn:  []int{3, 4},
				served:  []int{3, 4, 1, 2},
			},
			true,
		}, {
			"incorrect order - take out",
			args{
				takeOut: []int{1, 2},
				dineIn:  []int{3, 4},
				served:  []int{2, 1, 3, 4},
			},
			false,
		}, {
			"incorrect order - dine in",
			args{
				takeOut: []int{1, 2},
				dineIn:  []int{3, 4},
				served:  []int{1, 4, 2, 3},
			},
			false,
		}, {
			"given answer 1",
			args{
				takeOut: []int{1, 3, 5},
				dineIn:  []int{2, 4, 6},
				served:  []int{1, 2, 4, 6, 5, 3},
			},
			false,
		}, {
			"given answer 2",
			args{
				takeOut: []int{17, 8, 24},
				dineIn:  []int{12, 19, 2},
				served:  []int{17, 8, 12, 19, 24, 2},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFIFO(tt.args.takeOut, tt.args.dineIn, tt.args.served); got != tt.want {
				t.Errorf("isFIFO() = %v, want %v", got, tt.want)
			}
		})
	}
}

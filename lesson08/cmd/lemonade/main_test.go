package main

import "testing"

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		// TODO: Add test cases.
		{name: "ok", args: args{[]int{5, 10, 5, 20}}, want: true},
		{name: "not ok", args: args{[]int{10, 10}}, want: false},
		{name: "three by five - ok", args: args{
			[]int{5, 5, 5, 20},
		}, want: true},
		{name: "not a multiple of five", args: args{
			[]int{5, 10, 21},
		}, want: false},
		{name: "bills empty", args: args{[]int{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

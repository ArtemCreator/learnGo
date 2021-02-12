package main

import "testing"

func Test_mirrorTicket(t *testing.T) {
	type args struct {
		countTickets int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "ok", args: args{999}, want: 900},
		{name: "ok2", args: args{450}, want: 351},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorTicket(tt.args.countTickets); got != tt.want {
				t.Errorf("mirrorTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}
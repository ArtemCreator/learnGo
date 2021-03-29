package main

import "testing"

func Test_printDataType(t *testing.T) {
	type args struct {
		total int64
	}
	tests := []struct {
		name         string
		args         args
		wantTypeData string
	}{
		// TODO: Add test cases.
		{name: "1 * 1", args: args{1}, wantTypeData: "uint8"},
		{name: "1 * -1", args: args{-1}, wantTypeData: "int8"},
		{name: "640 * 100", args: args{64000}, wantTypeData: "uint16"},
		{name: "640 * -100", args: args{-64000}, wantTypeData: "int32"},
		{name: "3000 * 3000", args: args{9_000_000}, wantTypeData: "uint32"},
		{name: "3000 * -3000", args: args{-9_000_000}, wantTypeData: "int32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTypeData := printDataType(tt.args.total); gotTypeData != tt.wantTypeData {
				t.Errorf("printDataType() = %v, want %v", gotTypeData, tt.wantTypeData)
			}
		})
	}
}
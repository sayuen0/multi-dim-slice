package main

import (
	"reflect"
	"testing"
)

func Test_makeSlice(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{n: 101}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}},
	}
	for _, tt := range tests {
		if got := makeSlice(tt.args.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("makeSlice() = %v, want %v", got, tt.want)
		}
	}
}

func Test_to2D(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{args: args{s: makeSlice(20), n: 3}, want: [][]int{[]int{0, 1, 2}, []int{3, 4, 5}, []int{6, 7, 8}, []int{9, 10, 11}, []int{12, 13, 14}, []int{15, 16, 17}, []int{18, 19}}},
		{args: args{s: makeSlice(3), n: 1}, want: [][]int{[]int{0}, []int{1}, []int{2}}},
		{args: args{s: makeSlice(3), n: 3}, want: [][]int{[]int{0, 1, 2}}},
	}
	for _, tt := range tests {
		if got := to2D(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("to2D() = %v, want %v", got, tt.want)
		}
	}
}

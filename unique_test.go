package main

import (
	"testing"
)

func Test_unique1(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test (1,1,2,3,3)", args: args{[]int{1,1,2,3,3}}, want:2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique1(tt.args.a); got != tt.want {
				t.Errorf("unique1() = %v, want %v", got, tt.want)
			}
		})
	}
}
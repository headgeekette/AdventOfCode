package main

import (
	"testing"
)

func TestCompute(t *testing.T) {

	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	}

	for _, tt := range tests {
		got := Compute(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("Compute() = %v, want %v", got, tt.want)
		}
		for i, v := range got {
			if v != tt.want[i] {
				t.Errorf("Compute() = %v, want %v", got, tt.want)
			}
		}

	}
}

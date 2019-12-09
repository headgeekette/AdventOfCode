package main

import (
	"reflect"
	"testing"
)

func TestGetCoor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{"test1", "R75", []int{75, 0}},
		{"test2", "L75", []int{-75, 0}},
		{"test2", "U75", []int{0, 75}},
		{"test2", "D75", []int{0, -75}},
	}

	for _, tt := range tests {
		got := GetCoor(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("%s ::: GetCoor() = %v, want %v", tt.name, got, tt.want)
		}

		for i, v := range got {
			if v != tt.want[i] {
				t.Errorf("%s ::: GetCoor() = %v, want %v", tt.name, got, tt.want)
			}
		}
	}

}

func TestGetPath(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		want  [][]int
	}{
		{"Test 1", []string{"R8", "U5", "L5", "D3"}, [][]int{{0, 0}, {8, 0}, {8, 5}, {3, 5}, {3, 2}}},
		{"Test 2", []string{"U7", "R6", "D4", "L4"}, [][]int{{0, 0}, {0, 7}, {6, 7}, {6, 3}, {2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPath(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

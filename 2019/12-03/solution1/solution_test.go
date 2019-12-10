package main

import (
	"reflect"
	"testing"
)

func TestGetCoor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Point
	}{
		{"test1", "R75", Point{75, 0}},
		{"test2", "L75", Point{-75, 0}},
		{"test2", "U75", Point{0, 75}},
		{"test2", "D75", Point{0, -75}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCoor(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPath() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestGetPath(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		want  []Point
	}{
		{"Test 1", []string{"R8", "U5", "L5", "D3"}, []Point{Point{0, 0}, Point{8, 0}, Point{8, 5}, Point{3, 5}, Point{3, 2}}},
		{"Test 2", []string{"U7", "R6", "D4", "L4"}, []Point{Point{0, 0}, Point{0, 7}, Point{6, 7}, Point{6, 3}, Point{2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPath(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoCross(t *testing.T) {

	type segments struct {
		p1 Point
		q1 Point
		p2 Point
		q2 Point
	}

	tests := []struct {
		name  string
		input segments
		want  bool
	}{
		{"test 1", segments{Point{8, 0}, Point{8, 5}, Point{0, 7}, Point{6, 7}}, false},
		{"test 2", segments{Point{8, 5}, Point{3, 5}, Point{6, 7}, Point{6, 3}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DoCross(tt.input.p1, tt.input.q1, tt.input.p2, tt.input.q2) != tt.want {
				t.Errorf("DoCross() want %v", tt.want)
			}
		})
	}
}

func TestGetCross(t *testing.T) {

	type segments struct {
		p1 Point
		q1 Point
		p2 Point
		q2 Point
	}

	tests := []struct {
		name  string
		input segments
		want  Point
	}{
		{"test 1", segments{Point{8, 5}, Point{3, 5}, Point{6, 7}, Point{6, 3}}, Point{6, 5}},
		{"test 2", segments{Point{3, 5}, Point{3, 2}, Point{6, 3}, Point{2, 3}}, Point{3, 3}},
	}

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if DoCross(tt.input.p1, tt.input.q1, tt.input.p2, tt.input.q2) != tt.want {
	// 			t.Errorf("Cross() want %v", tt.want)
	// 		}
	// 	})
	// }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCross(tt.input.p1, tt.input.q1, tt.input.p2, tt.input.q2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCross() = %v, want %v", got, tt.want)
			}
		})
	}

}

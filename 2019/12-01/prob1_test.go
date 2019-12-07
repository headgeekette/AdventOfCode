package main

import (
	"testing"
)

func TestFuel(t *testing.T) {
	// type args struct {
	// 	mass int
	// }
	tests := []struct {
		input int
		want  int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tt := range tests {
		if got := Fuel(tt.input); got != tt.want {
			t.Errorf("Fuel() = %v, want %v", got, tt.want)
		}
	}
}

func TestExtFuel(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		if got := ExtFuel(tt.input); got != tt.want {
			t.Errorf("ExtFuel() = %v, want %v", got, tt.want)
		}
	}
}

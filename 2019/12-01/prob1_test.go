package main

import "testing"

func TestFuel(t *testing.T) {

	testvals := []struct {
		input    int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, testval := range testvals {
		result := Fuel(testval.input)
		if result != testval.expected {
			t.Errorf("Got %d, Want %d", result, testval.expected)
		}
	}
}

func TestExtFuel(t *testing.T) {

	testvals := []struct {
		input    int
		expected int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, testval := range testvals {
		initFuel := Fuel(testval.input)
		result := ExtFuel(initFuel)
		if result != testval.expected {
			t.Errorf("Got %d, Want %d", result, testval.expected)
		}
	}
}

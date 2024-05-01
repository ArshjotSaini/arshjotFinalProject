package utils

import (
	"image/color"
	"testing"
)

func TestGetStateFromLocationApi(t *testing.T) {
	cases := []struct {
		input          string
		expectedOutput string
	}{
		{"Salem, MA", "Salem, MA"},
		{"Atlanta, GA", "Atlanta, GA "},
		{"Jacksonville, FL", " Jacksonville, FL"},
		{"Los Angeles", "Los Angeles"},
	}

	for _, c := range cases {
		if output := GetStateFromLocationApi(c.input); output != c.expectedOutput {
			t.Errorf("incorrect output for `%s`: expected `%s` but got `%s`", c.input, c.expectedOutput, output)
		}
	}
}

func TestIsFileExist(t *testing.T) {
	cases := []struct {
		input          string
		expectedOutput bool
	}{
		{"handlers.go", true},
		{"maps.go", true},
		{"", true},
		{"helper.go", true},
		{"arshjot_test.go", true},
		{"./db/arshjot_final_project.db", true},
	}

	for _, c := range cases {
		if output := IsFileExist(c.input); output != c.expectedOutput {
			t.Errorf("incorrect output for `%v`: expected `%v` but got `%v`", c.input, c.expectedOutput, output)
		}
	}
}

func TestGetColor(t *testing.T) {
	cases := []struct {
		input          int
		expectedOutput color.RGBA
	}{
		{input: 0, expectedOutput: color.RGBA{R: 0x10, G: 0x10, B: 0x10, A: 0xff}},
		{input: 55, expectedOutput: color.RGBA{R: 0, G: 0xff, B: 0xff, A: 0xff}},
		{input: 23, expectedOutput: color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff}},
	}

	for _, c := range cases {
		if output := gotColor(c.input); output != c.expectedOutput {
			t.Errorf("incorrect output for `%v`: expected `%v` but got `%v`", c.input, c.expectedOutput, output)
		}
	}
}

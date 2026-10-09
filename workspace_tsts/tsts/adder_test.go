package tsts

import "testing"

func TestAddNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}

func TestAddTable(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"2+3=5", 2, 3, 5},
		{"0+0=0", 0, 0, 0},
		{"-1+1=0", -1, 1, 0},
		{"-1+6=0", -1, 6, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // В Go 1.27 этого достаточно, всё отработает корректно!
			result := addNumbers(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("%s: got %d, want %d", tt.name, result, tt.expected)
			}
		})
	}
}

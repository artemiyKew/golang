package test

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd_Bad(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(%d, %d) = %d; expected %d", 1, 2, result, 3)
	}
	result = Add(0, 0)
	if result != 0 {
		t.Errorf("Add(%d, %d) = %d; expected %d", 0, 0, result, 0)
	}
	result = Add(-1, -1)
	if result != -2 {
		t.Errorf("Add(%d, %d) = %d; expected %d", -1, -1, result, -2)
	}
}

// Table tests

func TestAdd_Good(t *testing.T) {
	tests := []struct {
		name           string
		a, b, expected int
	}{
		{name: "positive", a: 1, b: 2, expected: 3},
		{name: "negative", a: -1, b: -1, expected: -2},
		{name: "zero", a: 0, b: 0, expected: 0},
		{name: "positive and negative", a: -1, b: 1, expected: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
			}
		})
	}
}

// package main

// import "testing"

//	func TestAdd(t *testing.T) {
//		result := Add(2, 3)
//		expected := 5
//		if result != expected {
//			t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
//		}
//	}
package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	result = Add(-5, 10)
	expected = 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	result = Add(0, 0)
	expected = 0
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

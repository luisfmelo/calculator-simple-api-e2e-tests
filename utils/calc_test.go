// +build unit

package utils

import "testing"

func TestSum(t *testing.T) {
	if Sum(1,1) != 2 {
		t.Error("Sum is failing")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(2,1) != 1 {
		t.Error("Subtract is failing")
	}
}
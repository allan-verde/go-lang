package main

import (
	"testing"
)

func TestShouldAddCorrect(t *testing.T) {
	result := Add(5, 3)
	expected := 8.0

	if result != expected {
		t.Errorf("Incorrect result. Expected: %v, got: %v", expected, result)
	}
}

func TestShouldAddWithMultipleNumbersCorrect(t *testing.T) {
	result := Add(2, 4, 6, 8)
	expected := 20.0

	if result != expected {
		t.Errorf("Incorrect result. Expected: %v, got: %v", expected, result)
	}
}

func TestShouldSubtractCorrect(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2.0

	if result != expected {
		t.Errorf("Incorrect result. Expected: %v, got: %v", expected, result)
	}
}

func TestShouldMultiplyCorrect(t *testing.T) {
	result := Multiply(5, 3)
	expected := 15.0

	if result != expected {
		t.Errorf("Incorrect result. Expected: %v, got: %v", expected, result)
	}
}

func TestShouldMultiplyWithMultipleNumbersCorrect(t *testing.T) {
	result := Multiply(2, 4, 6, 8)
	expected := 384.0

	if result != expected {
		t.Errorf("Incorrect result. Expected: %v, got: %v", expected, result)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	result, err := Divide(10, 2)
	expected := 5.0

	if result != expected || err != nil {
		t.Errorf("Incorrect result. Expected: %v, got: %v, error: %v", expected, result, err)
	}
}

func TestShouldDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)

	if err == nil {
		t.Error("Expected error for division by zero")
	}
}

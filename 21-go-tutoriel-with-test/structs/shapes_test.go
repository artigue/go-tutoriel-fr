package main

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 5}
	expected := math.Pi * 25
	actual := c.Area()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	expected := 50.0
	actual := r.Area()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

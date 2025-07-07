package grade_calculator

import (
	"math"
	"testing"
)

func TestCalculateAverage_StandardGrades(t *testing.T) {
	gc := &GradeCalculator{}
	grades := []float64{90, 85, 88}
	expected := 87.66666666666667
	result := gc.CalculateAverage(grades)
	if math.Abs(result-expected) > 1e-9 {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculateAverage_EmptyList(t *testing.T) {
	gc := &GradeCalculator{}
	result := gc.CalculateAverage([]float64{})
	if result != 0.0 {
		t.Errorf("Expected 0.0, got %v", result)
	}
}

func TestCalculateAverage_NegativeGrades(t *testing.T) {
	gc := &GradeCalculator{}
	grades := []float64{100, -50}
	expected := 25.0
	result := gc.CalculateAverage(grades)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculateAverage_FractionalGrades(t *testing.T) {
	gc := &GradeCalculator{}
	grades := []float64{1.5, 2.5, 3.5}
	expected := 2.5
	result := gc.CalculateAverage(grades)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculateAverage_ExtremeValues(t *testing.T) {
	gc := &GradeCalculator{}
	grades := []float64{math.MaxFloat64, math.MaxFloat64}
	expected := math.Inf(1)
	result := gc.CalculateAverage(grades)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

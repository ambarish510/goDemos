package SamplePackage

import "testing"

func TestAvg(t *testing.T) {
	avgFound := FindAvg(2.5,3.5)
	if avgFound != float32(3) {
		t.Errorf("FindAvg(2.5,3.5) = %f; want 3.0", avgFound)
	}
}

func TestAvgNegatives(t *testing.T) {
	avgFound := FindAvg(-2.5,-3.5)
	if avgFound != float32(-3) {
		t.Errorf("FindAvg(-2.5,-3.5) = %f; want -3.0", avgFound)
	}
}

func TestAvgMixed(t *testing.T) {
	avgFound := FindAvg(-2.5,3.5)
	if avgFound != float32(0.5) {
		t.Errorf("FindAvg(-2.5,3.5) = %f; want 0.5", avgFound)
	}
}
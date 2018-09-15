package mymath

import "testing"

func TestSqrt(t *testing.T) {
	var tests = []struct {
		x      float64
		expect float64
	}{
		{4.0, 2.0},
	}

	for _, tCase := range tests {
		result := Sqrt(tCase.x)
		if result != tCase.expect {
			t.Errorf("Sqrt(%f): expect %f, actual %f", tCase.x, tCase.expect, result)
		}
	}
}

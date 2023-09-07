package mymath

import "testing"

type testData struct {
	x    float64
	y    float64
	want float64
}

func TestSqrt(t *testing.T) {
	testCases := []testData{
		{x: 4.0, want: 2.0},
		{x: 0, want: 0},
		{x: 1, want: 1},
		{x: 16.0, want: 4.0},
	}

	for _, tc := range testCases {
		res := Sqrt(tc.x)

		if res != tc.want {
			t.Errorf("Unexpected result. Got: %v, Want: %v", res, tc.want)
		}
	}
}

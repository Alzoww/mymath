package mymath

import "testing"

type testData1 struct {
	x    float64
	y    float64
	want float64
}

func TestSqrt(t *testing.T) {
	testCases := []testData1{
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

func TestCeil(t *testing.T) {
	testCases := []testData1{
		{x: 4.0, want: 4.0},
		{x: 5.345, want: 6.0},
		{x: 0.0, want: 0.0},
		{x: 4.5, want: 5.0},
	}

	for _, tc := range testCases {
		result := Ceil(tc.x)

		if result != tc.want {
			t.Errorf("Unexpected result. Got: %v, want: %v", result, tc.want)
		}
	}
}

func TestFloor(t *testing.T) {
	testCases := []testData1{
		{x: 4.0, want: 4.0},
		{x: 5.345, want: 5.0},
		{x: 0.0, want: 0.0},
		{x: 4.5, want: 4.0},
	}

	for _, tc := range testCases {
		result := Floor(tc.x)

		if result != tc.want {
			t.Errorf("Unexpected result. Got: %v, want: %v", result, tc.want)
		}
	}
}

func TestPow(t *testing.T) {
	testCases := []testData1{
		{x: 4.0, y: 2.0, want: 16.0},
		{x: 4.0, y: 0.0, want: 1.0},
		{x: 0, y: 2.0, want: 0.0},
		{x: 2.0, y: 3.0, want: 8.0},
	}

	for _, tc := range testCases {
		result := Pow(tc.x, tc.y)

		if result != tc.want {
			t.Errorf("Unexpected result. Got: %v, Want: %v", result, tc.want)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []testData1{
		{x: 2.5, y: 5.5, want: 5.5},
		{x: 2.5, y: 0.0, want: 2.5},
		{x: -1.5, y: 1.5, want: 1.5},
		{x: 2.5, y: 2.5, want: 2.5},
	}

	for _, tc := range testCases {
		result := Max(tc.x, tc.y)

		if result != tc.want {
			t.Errorf("Unexpected result. Got: %v, Want: %v", result, tc.want)
		}
	}
}

func TestMin(t *testing.T) {
	testCases := []testData1{
		{x: 2.5, y: 5.5, want: 2.5},
		{x: 2.5, y: 0.0, want: 0.0},
		{x: -1.5, y: 1.5, want: -1.5},
		{x: 2.5, y: 2.5, want: 2.5},
	}

	for _, tc := range testCases {
		result := Min(tc.x, tc.y)

		if result != tc.want {
			t.Errorf("Unexpected result. Got: %v, Want: %v", result, tc.want)
		}
	}
}

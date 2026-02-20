package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Testing perimeter of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		assertFloatEqual(t, got, want)
	})

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{5.0, 7.0}, 35.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{12.0, 6.0}, 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			assertFloatEqual(t, got, test.want)
		})
	}

}

func assertFloatEqual(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

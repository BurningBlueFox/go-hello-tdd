package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	check := func(t testing.TB, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("perimeter calculation of rectangle", func(t *testing.T) {
		got := Rectangle{10.0, 10.0}.Perimeter()
		want := 40.0
		check(t, got, want)
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt, got, tt.want)
		}
	}
}

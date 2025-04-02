package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("check perimeter of rectangle", func(t *testing.T) {
		got := Perimeter(Rectangle{10.0, 20.0})
		want := 60.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	}
// 	t.Run("check area of rectangle", func(t *testing.T) {
// 		rect := Rectangle{10.0, 20.0}
// 		// got := rect.Area()
// 		want := 200.0

// 		checkArea(t, rect, want)
// 	})

// 	t.Run("test area for cirlce", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		// got := circle.Area()
// 		want := 314.1592653589793

// 		checkArea(t, circle, want)
// 	})
// }

// Table driven test

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 20.0}, want: 200.0},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36},
	}

	for _, tt := range areaTest {
		// got := tt.shape.Area()
		// if got != tt.want {
		// 	t.Errorf("got %g want %g", got, tt.want)
		// }
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}

package main

import "testing"

const (
	shapeRectangle = "прямоугольник"
	shapeCircle    = "окружность"
)

func TestAreaCalculator(t *testing.T) {
	tests := []struct {
		name    string
		figures []Shape
		want    string
		want1   float64
	}{
		{
			"тест квадрат-квадрат",
			[]Shape{NewRectangle(1, 1, shapeRectangle), NewRectangle(2, 2, shapeRectangle)},
			"прямоугольник-прямоугольник",
			5,
		},
		{
			"тест квадрат-круг-квадрат",
			[]Shape{NewRectangle(1, 1, shapeRectangle), NewCircle(1, shapeCircle), NewRectangle(2, 2, shapeRectangle)},
			"прямоугольник-окружность-прямоугольник",
			8.14159,
		},
		{
			"пустота",
			nil,
			"",
			0,
		},
		{
			"тест олимпиада",
			[]Shape{NewCircle(1, shapeCircle), NewCircle(1, shapeCircle), NewCircle(1, shapeCircle), NewCircle(1, shapeCircle), NewCircle(1, shapeCircle)},
			"окружность-окружность-окружность-окружность-окружность",
			15.70795,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AreaCalculator(tt.figures)
			if got != tt.want {
				t.Errorf("AreaCalculator() got = %v, want %v", got, tt.want)
			}

			if got1 != tt.want1 {
				t.Errorf("AreaCalculator() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

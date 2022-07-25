package main

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	// TODO: код писать здесь
}

func NewRectangle(float64, float64, string) *Rectangle {
	// TODO: код писать здесь
	return &Rectangle{}
}

// TODO: код писать здесь

type Circle struct {
	// TODO: код писать здесь
}

func NewCircle(float64, string) *Circle {
	// TODO: код писать здесь
	return &Circle{}
}

func AreaCalculator(figures []Shape) (string, float64) {
	// TODO: код писать здесь
	return "", 0.0
}

func main() {
}

package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками на плоскости.
// Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64)
// и конструктором. Расстояние рассчитывается по формуле между координатами двух точек.
// Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.

// Point - структура с приватными полями
type Point struct {
	x float64
	y float64
}

// NewPoint - конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance - метод для расчета расстояния до другой точки
func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки: A и B
	pointA := NewPoint(1.5, 2.7)
	pointB := NewPoint(4.2, 5.3)

	// Вычисляем расстояние между A и B
	distance := pointA.Distance(pointB)

	// Выводим результат
	fmt.Printf("Расстояние между точками A(%.1f, %.1f) и B(%.1f, %.1f) = %.2f\n",
		pointA.x, pointA.y, pointB.x, pointB.y, distance)
	// Расстояние между точками A(1.5, 2.7) и B(4.2, 5.3) = 3.75

	// Создаем две другие точки: C и D
	pointC := NewPoint(-1, 0)
	pointD := NewPoint(0, 1)

	// Вычисляем расстояние между C и D
	distanceCD := pointC.Distance(pointD)

	// Выводим результат
	fmt.Printf("Расстояние между точками C(%.1f, %.1f) и D(%.1f, %.1f) = %.2f\n",
		pointC.x, pointC.y, pointD.x, pointD.y, distanceCD)
	// Расстояние между точками C(-1.0, 0.0) и D(0.0, 1.0) = 1.41
}

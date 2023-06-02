package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Square struct {
	Height float64
}

func (s Square) GetArea() float64 {
	return s.Height * s.Height
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) GetArea() float64 {
	return t.Base * t.Height / 2
}

// アクセス可能なメソッドを定義する
type Figure interface {
	GetArea() float64
}

//Figureの面積を表示する関数を定義 *4
func DisplayArea(f Figure) {
	fmt.Printf("面積は%vです\n", f.GetArea())
}

//実行される部分 *5
func ShapeFactory() {
	circle := Circle{Radius: 1.5}
	DisplayArea(circle)

	square := Square{Height: 3.0}
	DisplayArea(square)

	triangle := Triangle{Base: 4.0, Height: 3.0}
	DisplayArea(triangle)
}

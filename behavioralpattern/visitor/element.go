package main

// 抽象元素接口
type Shape interface {
	GetType() string
	Accept(Visitor) // 关键方法：接受一个访问者
}

// =========================================

// 具体元素类(圆形)
type Circle struct {
	Radius float64
}

func (c *Circle) GetType() string {
	return "Circle"
}
func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c) // 双分派：调用访问者中对应自己的方法
}

// 具体元素类(正方形)
type Square struct {
	Side float64
}

func (s *Square) GetType() string {
	return "Square"
}
func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s)
}

// 具体元素类(三角形)
type Triangle struct {
	Base, Height float64
}

func (t *Triangle) GetType() string {
	return "Triangle"
}
func (t *Triangle) Accept(v Visitor) {
	v.VisitTriangle(t)
}

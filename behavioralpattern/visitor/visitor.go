package main

import "fmt"

// 抽象访问者接口
type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Square)
	VisitTriangle(*Triangle)
}

// ====================================

// 具体访问者类：计算面积
type CalculateAreaVisitor struct {
	TotalArea float64
}

func (a *CalculateAreaVisitor) VisitCircle(c *Circle) {
	area := 3.14 * c.Radius * c.Radius
	fmt.Printf("Calculating area for circle (radius=%.2f): %.2f\n", c.Radius, area)
	a.TotalArea += area
}
func (a *CalculateAreaVisitor) VisitSquare(s *Square) {
	area := s.Side * s.Side
	fmt.Printf("Calculating area for square (side=%.2f): %.2f\n", s.Side, area)
	a.TotalArea += area
}
func (a *CalculateAreaVisitor) VisitTriangle(t *Triangle) {
	area := 0.5 * t.Base * t.Height
	fmt.Printf("Calculating area for triangle (base=%.2f, height=%.2f): %.2f\n", t.Base, t.Height, area)
	a.TotalArea += area
}

// 具体访问者类：导出XML
type ExportXMLVisitor struct{}

func (x *ExportXMLVisitor) VisitCircle(c *Circle) {
	fmt.Printf(`<Circle radius="%.2f"/>`+"\n", c.Radius)
}
func (x *ExportXMLVisitor) VisitSquare(s *Square) {
	fmt.Printf(`<Square side="%.2f"/>`+"\n", s.Side)
}
func (x *ExportXMLVisitor) VisitTriangle(t *Triangle) {
	fmt.Printf(`<Triangle base="%.2f" height="%.2f"/>`+"\n", t.Base, t.Height)
}
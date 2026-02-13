package main

import "fmt"

func main() {
	// 我们的图形结构
	shapes := []Shape{
		&Circle{Radius: 5},
		&Square{Side: 4},
		&Triangle{Base: 3, Height: 6},
	}

	// 操作1：计算总面积
	areaCalc := &CalculateAreaVisitor{}
	for _, shape := range shapes {
		shape.Accept(areaCalc) // 关键调用
	}
	fmt.Printf("Total area: %.2f\n\n", areaCalc.TotalArea)

	// 操作2：导出为XML
	xmlExporter := &ExportXMLVisitor{}
	fmt.Println("Exporting to XML:")
	for _, shape := range shapes {
		shape.Accept(xmlExporter)
	}
}

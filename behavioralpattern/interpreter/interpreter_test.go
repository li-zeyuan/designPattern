package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpret(t *testing.T)  {
	p := new(Parser)
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := p.Result().Interpret()
	fmt.Print(res)
}
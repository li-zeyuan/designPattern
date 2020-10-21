package decorator

import "testing"

func TestDecorator(t *testing.T)  {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)

	ret := c.Calc()
	t.Log(ret)
}

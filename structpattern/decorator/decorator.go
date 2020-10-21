package decorator

// 抽象构件
type Component interface {
	Calc() int
}

// 具体构件（被装饰的对象）
type ConcreteComponent struct {

}

func (c *ConcreteComponent)Calc() int  {
	return 0
}

// 抽象装饰（本例省略）

// 具体装饰角色(乘法）
type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component:c,
		num: num,
	}
}

func (m *MulDecorator)Calc()int {
	return m.Component.Calc() * m.num
}

// 具体装饰角色(加法）
type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component:c,
		num: num,
	}
}

func (a *AddDecorator)Calc()int  {
	return a.Component.Calc() + a.num
}
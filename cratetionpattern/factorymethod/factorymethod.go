package factorymethod

/*
基类->实际类->工厂类
 */

// -------------子类接口------------

// 封装子类的接口
type Operator interface {
	SetA(i int)
	SetB(i int)
	Result() int
}

// -------------基类---------------

// 抽离公用的属性和方法成 基类
type OperatorBase struct {
	A int
	B int
}

func (oBase *OperatorBase) SetA(a int) {
	oBase.A = a
}

func (oBase *OperatorBase) SetB(b int) {
	oBase.B = b
}

// ------------子类--------------

// 实际加法类
type PlusOperator struct {
	*OperatorBase
}

func (PlusO PlusOperator) Result() int {
	return PlusO.A + PlusO.B
}

// ------------工厂接口------------
type OperatorFactory interface {
	Create() Operator
}

// ------------工厂类--------------
// 加法工厂类
type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase:&OperatorBase{},
	}

}



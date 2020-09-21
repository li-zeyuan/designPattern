package adapter

// ========= 目标接口：客户按照这个接口访问适配者 ====

type Target interface {
	Request() string
}

// ========= 适配者：被访问的现存组件的接口=====

type Adaptee interface {
	SpecificReq() string
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct {
}

func (*adapteeImpl) SpecificReq() string {
	return "adaptee method"
}

// ============= 适配器：起转换作用 ========

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificReq()
}

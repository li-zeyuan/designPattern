package protorype

// ------------抽象原型-----------
type Cloneable interface {
	Clone() Cloneable
}

// ------------访问类------------
type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes:make(map[string]Cloneable)}
}

func (p *PrototypeManager)Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager)Set(name string, prototype Cloneable)  {
	p.prototypes[name] = prototype
}
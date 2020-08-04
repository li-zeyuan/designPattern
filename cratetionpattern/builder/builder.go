package builder

// --------------抽象构建者-------------
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// --------------指挥者-----------------
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder:builder,
	}
}

func (dir *Director) Construct() {
	dir.builder.Part1()
	dir.builder.Part2()
	dir.builder.Part3()
}

// ---------------产品角色----------------
type Product1 struct {
	result string
}

func (b *Product1) Part1() {
	b.result += "1"
}

func (b *Product1) Part2() {
	b.result += "2"
}

func (b *Product1) Part3()  {
	b.result += "3"
}

func (b *Product1) GetResult() string {
	return b.result
}

type Product2 struct {
	result int
}

func (b *Product2) Part1() {
	b.result += 1
}

func (b *Product2) Part2()  {
	b.result += 2
}

func (b *Product2) Part3()  {
	b.result += 3
}

func (b *Product2) GetResult() int {
	return b.result
}
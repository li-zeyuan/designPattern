package iterator

// 抽象聚合接口
type Aggregate interface {
	Iterator() Iterator
}

// 具体聚合类
type Numbers struct {
	start int
	end   int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers)Iterator() Iterator {
	return &NumberIterator{
		numbers: n,
		next:    n.start,
	}
}

// 抽象迭代器接口
type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

// 具体迭代器类
type NumberIterator struct {
	numbers *Numbers
	next int
}

func (i *NumberIterator)First()  {
	i.next = i.numbers.start
}

func (i *NumberIterator)IsDone() bool {
	return i.next > i.numbers.end
}

func (i *NumberIterator)Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next ++
		return next
	}

	return nil
}
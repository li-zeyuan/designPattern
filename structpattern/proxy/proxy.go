package proxy

// 抽象主题
type Subject interface {
	Do() string
}

// 真正的对象
type ReadSubject struct {

}

func (r *ReadSubject)Do() string {
	return "real"
}

// 代理
type Proxy struct {
	read ReadSubject
}

func (p *Proxy) Do()string {
	result := ""
	// 代理前处理
	result += "pre:"

	result += p.read.Do()

	// 代理后处理
	result += ":after"
	return result
}




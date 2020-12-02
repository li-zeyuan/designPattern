package state

import "fmt"

// 环境类
type DayContext struct {
	today Week // 持有状态类接口
}

func NewDayContext() *DayContext {
	ctx := new(DayContext)
	ctx.today = new(Monday)

	return ctx
}

func (d *DayContext) Today() {
	d.today.Today()
}


func (d *DayContext) Next() {
	d.today.Next(d)
}

// 抽象状态接口
type Week interface {
	Today()
	Next(ctx *DayContext)
}

// 具体状态类
type Monday struct{}

func (*Monday) Today() {
	fmt.Printf("Monday\n")
}

func (*Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

type Tuesday struct{}

func (*Tuesday) Today() {
	fmt.Printf("Tuesday\n")
}

func (*Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

type Wednesday struct{}

func (*Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct{}

func (*Thursday) Today() {
	fmt.Printf("Thursday\n")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct{}

func (*Friday) Today() {
	fmt.Printf("Friday\n")
}

func (*Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct{}

func (*Saturday) Today() {
	fmt.Printf("Saturday\n")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}

type Sunday struct{}

func (*Sunday) Today() {
	fmt.Printf("Sunday\n")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}
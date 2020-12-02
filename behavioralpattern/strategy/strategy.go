package strategy

import "fmt"

type PaymentContext struct {
	Name   string
	CardId string
	Money  int
}

// 环境类
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy // 持有子类的引用
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	payment := new(Payment)

	ctx := new(PaymentContext)
	ctx.Name = name
	ctx.CardId = cardId
	ctx.Money = money

	payment.context = ctx
	payment.strategy = strategy

	return payment
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

// 抽象策略接口
type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

// 具体策略类1
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by cash: %s.", ctx.Money, ctx.Name, ctx.CardId)
}

// 具体策略类2
type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by bank: %s.", ctx.Money, ctx.Name, ctx.CardId)
}

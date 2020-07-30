package strategy

import "fmt"

// define one set of replaceable functions, different functions will be used in different scenarios
type PayStrategy interface {
	pay(ctx *PaymentContext)
}
type PaymentContext struct {
	name      string
	id, money int
}

type Payment struct {
	*PaymentContext
	PayStrategy
}

func (p *Payment) Pay() {
	p.PayStrategy.pay(p.PaymentContext)
}
func NewPayment(name string, id, money int, way PayStrategy) *Payment {
	return &Payment{
		PaymentContext: &PaymentContext{
			name:  name,
			id:    id,
			money: money,
		},
		PayStrategy: way,
	}
}

// cash pay
type Cash struct {
}

func (c *Cash) pay(ctx *PaymentContext) {
	fmt.Printf("cash pay - name=%s, card id=%d, money=%d\n", ctx.name, ctx.id, ctx.money)
}

// AliPay
type AliPay struct {
}

func (c *AliPay) pay(ctx *PaymentContext) {
	fmt.Printf("ali pay - name=%s, card id=%d, money=%d\n", ctx.name, ctx.id, ctx.money)
}

// WeChat pay
type WeChat struct {
}

func (c *WeChat) pay(ctx *PaymentContext) {
	fmt.Printf("wechat pay - name=%s, card id=%d, money=%d\n", ctx.name, ctx.id, ctx.money)
}

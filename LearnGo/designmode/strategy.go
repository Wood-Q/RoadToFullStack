package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount int) string
}
type Alipay struct{}

func (a Alipay) Pay(amount int) string {
	return fmt.Sprintf("Alipay %d", amount)
}

type WechatPay struct{}

func (w WechatPay) Pay(amount int) string {
	return fmt.Sprintf("Wechat Pay %d", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount int) string {
	return p.strategy.Pay(amount)
}

func main() {
	ctx := NewPaymentContext(WechatPay{})
	fmt.Println(ctx.Pay(30))
	ctx.SetStrategy(Alipay{})
	fmt.Println(ctx.Pay(30))
}

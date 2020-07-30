package strategy

import "testing"

func TestDemo(t *testing.T) {
	ali := NewPayment("kevin", 1000, 500, &AliPay{})
	ali.Pay()

	wecaht := NewPayment("kevin", 1000, 500, &WeChat{})
	wecaht.Pay()

	cash := NewPayment("kevin", 1000, 500, &Cash{})
	cash.Pay()
}

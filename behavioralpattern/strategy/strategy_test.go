package strategy

import "testing"

func TestCash_Pay(t *testing.T) {
	cash := new(Cash)
	payment := NewPayment("lizeyuan", "10086", 1, cash)
	payment.Pay()
}

func TestBank_Pay(t *testing.T) {
	bank := new(Bank)
	payment := NewPayment("lizeyuan", "10086", 1, bank)
	payment.Pay()
}

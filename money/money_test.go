package money

import "testing"

func TestMoneyMultplication(t *testing.T)  {
	five := NewDollar	(5)
	five.times(2)
	if five.amount != 10 {
		t.Errorf("want 10, got %d", five.amount)
	}
}

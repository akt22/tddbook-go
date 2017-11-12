package money

import (
	"testing"
)

func TestMoneyMultplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if product.amount != 10 {
		t.Errorf("want 10, got %d", product.amount)
	}

	// 副作用をテスト
	product = five.times(3)
	if product.amount != 15 {
		t.Errorf("want 15, got %d", product.amount)
	}
}

func TestEquality(t *testing.T) {
	d := Dollar{5}
	if !d.equals(Dollar{5}) {
		t.Errorf("want same, but different")
	}
	if d.equals(Dollar{6}) {
		t.Errorf("want false, got true")
	}
}

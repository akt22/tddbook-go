package money

import (
	"testing"
)

func TestMoneyMultplication(t *testing.T) {
	five := NewDollar(5)
	// 掛け算のテスト & 等価性のテストが通っているという仮定の基進んだ
	if NewDollar(10) == five.times(2) {
		t.Errorf("want 10, got %d", five.times(2))
	}

	// 副作用をテスト
	if NewDollar(15) == five.times(3) {
		t.Errorf("want 15, got %d", five.times(3))
	}
}

func TestEquality(t *testing.T) {
	d := NewDollar(5)
	if !d.equals(NewDollar(5)) {
		t.Errorf("want same, but different")
	}
	if d.equals(NewDollar(6)) {
		t.Errorf("want false, got true")
	}

	if NewFranc(5).equals(NewDollar(5)) {
		t.Errorf("want false, got true")
	}
}

func TestCurrency(t *testing.T) {
	if "USD" != NewDollar(1).Currency() {
		t.Errorf("want USD, got %s", NewDollar(1).Currency())
	}
	if "CHF" != NewFranc(1).Currency() {
		t.Errorf("want CHF, got %s", NewFranc(1).Currency())
	}
}

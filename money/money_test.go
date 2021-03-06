package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoneyMultplication(t *testing.T) {
	five := NewDollar(5)
	// 掛け算のテスト & 等価性のテストが通っているという仮定の基進んだ
	if NewDollar(10) == five.Times(2) {
		t.Errorf("want 10, got %d", five.Times(2))
	}

	// 副作用をテスト
	if NewDollar(15) == five.Times(3) {
		t.Errorf("want 15, got %d", five.Times(3))
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

func TestSimpleAddiction(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := Bank{}
	reduced := bank.reduce(sum, "USD")
	if !NewDollar(10).equals(reduced) {
		t.Errorf("want 10, got %d", reduced.Amount())
	}
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum, ok := result.(*Sum)
	if !ok {
		t.Fatalf("want Sum")
	}
	if !five.equals(sum.augend) {
		t.Errorf("want 5, got %d", sum.augend)
	}
	if !five.equals(sum.addend) {
		t.Errorf("want 5, got %d", sum.addend)
	}
}

func TestReduceSum(t *testing.T) {
	sum := &Sum{NewDollar(3), NewDollar(4)}
	bank := NewBank()
	result := bank.reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	result := bank.reduce(NewFranc(2), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestIdentifyRate(t *testing.T) {
	bank := NewBank()
	assert.Equal(t, 1, bank.rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	result := bank.reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, NewDollar(10), result)
}

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Plus(fiveBucks)
	result := bank.reduce(sum, "USD")
	assert.Equal(t, NewDollar(15), result)
}

func TestSumTimesMoney(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Times(2)
	result := bank.reduce(sum, "USD")
	assert.Equal(t, NewDollar(20), result)
}

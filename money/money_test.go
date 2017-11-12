package money

import (
	"testing"
)

func TestMoneyMultplication(t *testing.T) {
	five := NewDollar(5)
	if NewDollar(10) == five.times(2) {
		t.Errorf("want 10, got %d", five.times(2))
	}

	// 副作用をテスト
	if NewDollar(15) == five.times(3) {
		t.Errorf("want 15, got %d", five.times(3))
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

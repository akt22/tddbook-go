package money

// =========================
// Money
// -------------------------
type IMoney interface {
	Amount() int
}

type Money struct {
	amount int
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) equals(object interface{}) bool {
	money := object.(IMoney)
	return m.Amount() == money.Amount()
}


// =========================
// Doller
// -------------------------
type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		Money{amount},
	}
}

func (d *Dollar) times(multiplier int) *Dollar {
	return &Dollar{
		Money{d.Money.amount * multiplier},
	}
}

// =========================
// Franc
// -------------------------
type Franc struct {
	Money
}

func NewFranc(amount int) *Franc {
	return &Franc{
		Money{amount},
	}
}

func (f *Franc) times(multiplier int) *Franc {
	return &Franc{
		Money{f.amount * multiplier},
	}
}

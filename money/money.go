package money

// =========================
// Money
// -------------------------
type IMoney interface {
	Amount() int
	Currency() string
}

type Money struct {
	amount   int
	currency string
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) equals(object interface{}) bool {
	money := object.(IMoney)
	return m.Amount() == money.Amount() &&
		m.Currency() == money.Currency()
}

// =========================
// Doller
// -------------------------
type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		Money{
			amount:   amount,
			currency: "USD",
		},
	}
}

func (d *Dollar) times(multiplier int) IMoney {
	return &Dollar{
		Money{
			amount:   d.Money.amount * multiplier,
			currency: "USD",
		},
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
		Money{
			amount:   amount,
			currency: "CHF",
		},
	}
}

func (f *Franc) times(multiplier int) IMoney {
	return &Franc{
		Money{
			amount:   f.Money.amount * multiplier,
			currency: "CHF",
		},
	}
}

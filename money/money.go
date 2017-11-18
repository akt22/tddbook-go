package money

import "fmt"

// =========================
// Sum
// -------------------------
type Sum struct {
	augend, addend *Money
}

// =========================
// Bank
// -------------------------
type Bank struct{}

func (b *Bank) reduce(source Expression, to string) *Money {
	return NewDollar(10)
}

// =========================
// Expression
// -------------------------
type Expression interface{}

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
	mm := object.(IMoney)
	return m.amount == mm.Amount() &&
		m.currency == mm.Currency()
}

func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount(), m.Currency())
}

func (m *Money) times(multiplier int) *Money {
	return &Money{
		m.Amount() * multiplier,
		m.Currency(),
	}
}

func (m *Money) plus(added *Money) Expression {
	return &Sum{m, added}
}

// =========================
// Dollar
// -------------------------
func NewDollar(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "USD",
	}
}

// =========================
// Franc
// -------------------------
func NewFranc(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "CHF",
	}
}

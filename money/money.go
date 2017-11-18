package money

import "fmt"

// =========================
// Bank
// -------------------------
type Bank struct{}

func (b *Bank) reduce(source Expression, to string) *Money {
	return source.Reduce(to)
}

// =========================
// Expression
// -------------------------
type Expression interface {
	Reduce(to string) *Money
}

// =========================
// Sum
// -------------------------
type Sum struct {
	addend, augend *Money
}

func (s *Sum) Reduce(to string) *Money {
	amount := s.augend.amount + s.addend.amount
	return NewMoney(amount, to)
}

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

func NewMoney(amount int, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) equals(object interface{}) bool {
	mm := object.(IMoney)
	return m.amount == mm.Amount() &&
		m.currency == mm.Currency()
}

func (m *Money) Currency() string {
	return m.currency
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

func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount(), m.Currency())
}

func (m *Money) Reduce(to string) *Money {
	return m
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

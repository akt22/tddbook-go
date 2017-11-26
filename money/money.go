package money

import "fmt"

// =========================
// Bank
// -------------------------
type Bank struct{}

func (b *Bank) reduce(source Expression, to string) *Money {
	return source.Reduce(*b, to)
}

func (b *Bank) addRate(from string, to string, rate int) {

}

func (b *Bank) rate(from string, to string) int {
	if from == "CHF" && to == "USD" {
		return 2
	}
	return 1
}

// =========================
// Expression
// -------------------------
type Expression interface {
	Reduce(bank Bank, to string) *Money
}

// =========================
// Sum
// -------------------------
type Sum struct {
	addend, augend *Money
}

func (s *Sum) Reduce(bank Bank, to string) *Money {
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

func (m *Money) Reduce(bank Bank, to string) *Money {
	rate := bank.rate(m.currency, to)
	NewMoney(m.amount/rate, to)
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

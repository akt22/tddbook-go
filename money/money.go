package money

import "fmt"

// =========================
// Bank
// -------------------------
type Bank struct {
	rates map[Pair]int
}

func (b *Bank) reduce(source Expression, to string) *Money {
	return source.Reduce(*b, to)
}

func (b *Bank) addRate(from string, to string, rate int) {
	b.rates[Pair{from, to}] = rate
}

func (b *Bank) rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rates[Pair{from, to}]
}

func NewBank() *Bank {
	return &Bank{
		rates: make(map[Pair]int),
	}
}

// =========================
// Expression
// -------------------------
type Expression interface {
	Reduce(bank Bank, to string) *Money
	Plus(addend Expression) Expression
	Times(multiplier int) Expression
}

// =========================
// Sum
// -------------------------
type Sum struct {
	addend, augend Expression
}

func (s *Sum) Reduce(bank Bank, to string) *Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return NewMoney(amount, to)
}

func (s *Sum) Plus(addend Expression) Expression {
	return &Sum{s, addend}
}

func (s *Sum) Times(multiplier int) Expression {
	return &Sum{s.augend.Times(multiplier), s.addend.Times(multiplier)}
}

func NewSum(addend, augend Expression) *Sum {
	return &Sum{addend, augend}
}

// =========================
// Pair
// -------------------------
type Pair struct {
	from, to string
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
	return m.Amount() == mm.Amount() &&
		m.Currency() == mm.Currency()
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Times(multiplier int) Expression {
	return NewMoney(m.Amount()*multiplier, m.currency)
}

func (m *Money) Plus(addend Expression) Expression {
	return &Sum{m, addend}
}

func (m *Money) Reduce(bank Bank, to string) *Money {
	rate := bank.rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
}

func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.amount, m.currency)
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

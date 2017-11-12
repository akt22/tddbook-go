package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (d *Dollar) times(multiplier int) *Dollar {
	return &Dollar{
		amount: d.amount * multiplier,
	}
}

func (d *Dollar) equals(object interface{}) bool {
	dollar := object.(Dollar)
	return d.amount == dollar.amount
}


type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{
		amount: amount,
	}
}

func (f *Franc) times(multiplier int) *Franc {
	return &Franc{
		amount: f.amount * multiplier,
	}
}

func (f *Franc) equals(object interface{}) bool {
	franc := object.(Franc)
	return f.amount == franc.amount
}

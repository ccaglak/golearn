package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("can't withdraw, insufficient funds")

func (w *Wallet) WithDraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

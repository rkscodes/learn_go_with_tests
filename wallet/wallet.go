package wallet

import (
	"errors"
	"fmt"
)

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amt BitCoin) {
	w.balance += amt
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) WithDraw(amt BitCoin) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amt
	return nil
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

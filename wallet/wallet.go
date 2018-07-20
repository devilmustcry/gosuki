package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin : Bitconnect!!!!!
type Bitcoin int

// ErrInsufficientFunds : Error thrown when error insufficient
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Wallet : Wallet for keep money
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// type WalletBehavior interface {
// 	Deposit()
// 	Balance() int
// }

// Deposit : deposit money into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance : get balance into wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw : Remove money from wallet (Only when BTC Go MOON)
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

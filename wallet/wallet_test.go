package wallet

import (
	"testing"
)

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error bud didn't get one")
	}
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Wanted an error but didn't get one")
	}
}
func TestWallet(t *testing.T) {
	t.Run("Should be able to deposit the bitcoin balance into Wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Should be able to get Bitcoin in BTC currency", func(t *testing.T) {
		satoshi := Bitcoin(1)
		got := satoshi.String()
		want := "1 BTC"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("Should be able to withdraw from Wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
	t.Run("Should not be able to withdraw if not enough money", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

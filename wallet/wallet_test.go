package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		err := wallet.WithDraw(BitCoin(5))
		assertNoError(t, err)

		assertBalance(t, wallet, BitCoin(5))

	})

	t.Run("withdraw more than balance", func(t *testing.T) {
		startingBalance := BitCoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.WithDraw(BitCoin(50))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
		// assertBalance(t, wallet, BitCoin(5))

	})
}

func assertBalance(t *testing.T, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got.Error(), want.Error())
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("wanted no error, got %s", got.Error())
	}
}

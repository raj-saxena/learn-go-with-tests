package pointersAndErrors

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		amount := Bitcoin(10)
		wallet := Wallet{}

		wallet.Deposit(amount)

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}

		err := wallet.Withdraw(Bitcoin(30))
		want := Bitcoin(70)

		got := wallet.Balance()

		if err != nil {
			t.Errorf("Not cool, got error: %v", err)
		}
		assertBalance(t, got, want)
	})

	t.Run("Err on insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		got := wallet.Withdraw(Bitcoin(20))

		if got == nil {
			t.Fatal("We were actually hoping for an error. :( ")
		}

		want := ErrInsufficientFunds

		if got != want {
			t.Errorf("got '%v', want '%v'", got, want)
		}
	})
}

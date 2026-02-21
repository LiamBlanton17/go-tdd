package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertEqualBitcoin(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(25)}

		gotErr := wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(15)

		assertNoError(t, gotErr)
		assertEqualBitcoin(t, got, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(25)}

		gotErr := wallet.Withdraw(50)
		got := wallet.Balance()
		want := Bitcoin(25)

		assertEqualBitcoin(t, got, want)
		assertError(t, gotErr, ErrorInsufficientFunds)
	})

}

func assertEqualBitcoin(t testing.TB, got, want Bitcoin) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Wanted an error but did not get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("Got an error but did not want one")
	}
}

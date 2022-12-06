package bank_test

import (
	"github.com/rangzen/go-boataootz/example/bank"
	"github.com/rangzen/go-boataootz/test"
	"testing"
)

func TestBalance(t *testing.T) {
	transactions := []bank.Transaction{
		{
			From: "Bartolomeo",
			To:   "Charlie",
			Sum:  100,
		},
		{
			From: "Albert",
			To:   "Bartolomeo",
			Sum:  25,
		},
	}

	test.AssertEqual(t, bank.BalanceFor(transactions, "Albert"), -25)
	test.AssertEqual(t, bank.BalanceFor(transactions, "Bartolomeo"), -75)
	test.AssertEqual(t, bank.BalanceFor(transactions, "Charlie"), 100)

	test.AssertEqual(t, bank.BalanceForMap(transactions, "Albert"), -25)
	test.AssertEqual(t, bank.BalanceForMap(transactions, "Bartolomeo"), -75)
	test.AssertEqual(t, bank.BalanceForMap(transactions, "Charlie"), 100)
}

package bank

import f "github.com/rangzen/go-boataootz/functional"

type Transaction struct {
	From string
	To   string
	Sum  float64
}

// BalanceFor calculates balance for a user in a collection of Transaction.
// It uses an accumulator with logic.
func BalanceFor(transactions []Transaction, name string) float64 {
	initialValue := 0.
	accumulator := func(balance float64, t Transaction) float64 {
		switch name {
		case t.From:
			return balance - t.Sum
		case t.To:
			return balance + t.Sum
		default:
			return balance
		}
	}
	return f.Reduce(transactions, accumulator, initialValue)
}

// BalanceForMap calculates balance for a user in a collection of Transaction.
// It uses a map, then a simple accumulator.
func BalanceForMap(transactions []Transaction, name string) float64 {
	initialValue := 0.
	accumulator := func(a, b float64) float64 {
		return a + b
	}
	ts := f.Map(transactions, func(t Transaction) float64 {
		switch name {
		case t.From:
			return -1 * t.Sum
		case t.To:
			return t.Sum
		default:
			return 0
		}
	})
	return f.Reduce(ts, accumulator, initialValue)
}

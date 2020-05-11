package account

import (
	"sync"
)

type Account struct {
	balance int
	closed bool
	access sync.Mutex
}

func (acc *Account) Balance() (int, bool) {
	if acc.closed {
		return 0, false
	}
	return acc.balance, true
}

func (acc *Account) Close() (int, bool) {
	acc.access.Lock()
	if acc.closed {
		acc.access.Unlock()
		return 0, false
	}
	acc.closed = true
	acc.access.Unlock()
	return acc.balance, true
}

func (acc *Account) Deposit(amount int) (int, bool) {
	acc.access.Lock()
	if acc.closed {
		acc.access.Unlock()
		return 0, false
	}
	if (acc.balance + amount) < 0 {
		acc.access.Unlock()
		return acc.balance, false
	}
	acc.balance += amount
	acc.access.Unlock()
	return acc.balance, true
}

func Open(amount int) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount}
}


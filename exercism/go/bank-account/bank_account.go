package account

import "sync"

// Account type.
type Account struct {
	Bal  int64
	Open bool
	lock *sync.Mutex
}

// Open opens an account. On opening, its initial balance will not be
// negative. If a negative balance is given to the Open function, nil is
// returned, otherwise the newly created account is returned.
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{Bal: amount, Open: true, lock: &sync.Mutex{}}
}

// Balance returns the balance of the account and a boolean indicating if the
// operation succeeded. Checking the balance does not succeed if the account is
// closed.
func (a *Account) Balance() (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.Bal, a.Open
}

// Deposit handles both deposits and withdrawals. If the argument is positive,
// then that amount is deposited in the account. If the amount negative, amount
// is withdrawn from the account. If the account is closed, its balance is not
// modified. Deposit returns the new balance of the account and a boolean that
// indicates if the operation succeeded. Deposit fails if the account is closed
// or if there is not enough money to withdraw from the account.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.Open && a.Bal+amount >= 0 {
		a.Bal += amount
		return a.Bal, true
	}
	return 0, false
}

// Close closes an account. It returns the balance the account has and a boolean
// true indicating the account was closed successfully. Closing an account does
// not succeed if the account is already closed. When an account is closed, its
// balance must be set to 0.
func (a *Account) Close() (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.Open {
		a.Open = false
		amt := a.Bal
		a.Bal = 0
		return amt, true
	}
	return 0, false
}

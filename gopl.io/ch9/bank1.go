package ch9

import "sync"

var (
	sema = make(chan struct{}, 1)
	balance int
	mu sync.Mutex
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<- sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

func Withdraw(amount int) bool {
	sema <- struct{}{}
	b := balance
	defer func() {
		<- sema
	}()

	if amount <= b {
		balance -= amount

		return true
	}

	return false
}
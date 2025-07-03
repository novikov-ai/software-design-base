package main

import (
	"fmt"
	
	"software-design-base/internal/payment"
	"software-design-base/internal/user"
)

func main() {
	bankAccount := payment.New(250000)

	user := user.New(bankAccount)

	fmt.Printf("User init with balance: %v\n", user.Account.GetBalance())

	putMoney := 10000

	user.Account.Deposit(float64(putMoney))

	fmt.Printf("User popped up by %v and the balance is: %v\n", putMoney, user.Account.GetBalance())

	tookMoney := 300000

	user.Account.Withdraw(float64(tookMoney))
	fmt.Printf("User withdrawed %v and the balance is: %v\n", tookMoney, user.Account.GetBalance())
}

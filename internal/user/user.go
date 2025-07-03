package user

import "software-design-base/internal/payment"

type User struct {
	Account payment.Account
}

func New(account payment.Account) *User {
	return &User{
		Account: account,
	}
}

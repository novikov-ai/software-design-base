package payment

type Account interface {
	Deposit(float64)
	Withdraw(float64)
	GetBalance() float64
}

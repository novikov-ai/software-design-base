package payment

type BankAccount struct {
	balance float64
}

func New(balance float64) *BankAccount {
	return &BankAccount{
		balance: balance,
	}
}

func (ba *BankAccount) Deposit(sum float64) {
	ba.balance += sum
}

func (ba *BankAccount) Withdraw(amount float64) {
	ba.balance -= amount
}

func (ba *BankAccount) GetBalance() float64 {
	return ba.balance
}



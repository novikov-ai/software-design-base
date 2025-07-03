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
	if sum < 0 {
		return
	}

	ba.balance += sum
}

func (ba *BankAccount) Withdraw(amount float64) {
	if amount < 0 {
		println("WARNING: can't withdraw negative ammount")
		return
	}

	if amount > ba.balance{
		println("WARNING: no enough money to withdraw")
		return
	}

	ba.balance -= amount
}

func (ba *BankAccount) GetBalance() float64 {
	return ba.balance
}

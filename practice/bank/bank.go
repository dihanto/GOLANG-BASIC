package bank

type BankAccount struct {
	Name          string
	AccountNumber int
	Balance       int
}

func (account *BankAccount) Withdraw(ammount int) {
	account.Balance -= ammount
}

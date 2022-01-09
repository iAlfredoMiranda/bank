package accounts

import "bank/holders"

type SavingsAccount struct {
	AccountHolder                                   holders.Holder
	NumberAgency, NumberCont, DigitCont, NumberBank int
	BankName                                        string
	AccountBalance                                  float64
}

//func for realize validation of Deposit
func (c *SavingsAccount) Deposit(valueDeposit float64) string { //func recive value of Current Account, value of Ted and return a string

	if valueDeposit > 0 { //rule for validation of ted
		c.AccountBalance += valueDeposit //iF true accountBalance + valueTed

		return "Deposit performed successfully"

	} else {
		return ("Value for deposit unavailable")
	}

}

package holders

import "fmt"

type CurrentAccount struct { //Current account data "struct"
	AccountHolder  string
	NumberAgency   int
	NumberCont     int
	DigitCont      int
	NumberBank     int
	BankName       string
	AccountBalance float64
}

//---------------------accountHolder / numberAgency / numberCont / digitCont / numberBank / bankName / accountBalance
var accountAlfredo = CurrentAccount{"Alfredo", 001, 4760850, 5, 0260, "Nu Pagamentos S.A - Instituição Financeira", 1800} //Mode most indicated for use of struct with Golang, defined definido pela sequência de datas estruturais
var accountLarissa = CurrentAccount{"Larissa", 001, 4760859, 5, 0260, "Nu Pagamentos S.A - Instituição Financeira", 1000} //Mode most indicated for use of struct with Golang, defined definido pela sequência de datas estruturais

//func for operation TED
func opTransfer() {

	fmt.Println("What is the value for TED?")
	fmt.Println("Balance available is R$", +accountAlfredo.AccountBalance) //show account balanace fo account Alfredo
	fmt.Println("Digit the value for operation: ")

	valueOperationTed := readCmdTed() // read value for operation

	fmt.Scan(accountAlfredo.Transfer(valueOperationTed, &accountLarissa)) // call to realize operation transfer if validation is success

	fmt.Println("Your balance now is: R$", +accountAlfredo.AccountBalance) //show account balanace fo account Alfredo att
	fmt.Println("Larissa saldo: ", +accountLarissa.AccountBalance)

}

//func for realize validation of Ted
func (c *CurrentAccount) Ted(valueTed float64) string { //func recive value of Current Account, value of Ted and return a string

	canTed := valueTed > 0 && valueTed <= c.AccountBalance //rule for validation of ted

	if canTed {
		c.AccountBalance -= valueTed //iF true accountBalance - valueTed
		return "Transfer performed successfully"

	} else if valueTed <= 0 {
		return "Transfer unavailable"

	} else {
		return "Account balance insufficient"
	}

}

//func for realize Transfer
func (c *CurrentAccount) Transfer(valueTed float64, destinationAccount *CurrentAccount) bool {
	if valueTed < c.AccountBalance && valueTed > 0 {
		c.AccountBalance -= valueTed
		destinationAccount.Deposit(valueTed)
		fmt.Println("Transfer performed successfully")
		return true
	} else if valueTed <= 0 {
		fmt.Println("Transfer unavailable")
		return false
	} else {
		fmt.Println("Account balance insufficient")
		return false
	}

}

//func for realize Deposit
func opDeposit() {

	fmt.Println("What is the value for Deposit?")
	fmt.Println("Your Balance available is R$", +accountAlfredo.AccountBalance)

	valueOperationDeposit := readCmdDeposit() // read value for operation

	fmt.Println(accountAlfredo.Deposit((valueOperationDeposit))) // call to realize operation deposit if validation is success
	fmt.Println("Your balance now is: R$", +accountAlfredo.AccountBalance)

}

//func for realize validation of Deposit
func (c *CurrentAccount) Deposit(valueDeposit float64) string { //func recive value of Current Account, value of Ted and return a string

	if valueDeposit > 0 { //rule for validation of ted
		c.AccountBalance += valueDeposit //iF true accountBalance + valueTed

		return "Deposit performed successfully"

	} else {
		return ("Value for deposit unavailable")
	}

}

//func for read operation Ted
func readCmdTed() float64 {

	var readCmdTed float64 //var for read user command

	fmt.Scan(&readCmdTed)
	fmt.Println("Value for Ted is: ", readCmdTed)

	return readCmdTed

}

//func for read operation Deposit
func readCmdDeposit() float64 {

	var readCmdDeposit float64 //var for read user command

	fmt.Scan(&readCmdDeposit)                             //recive user command
	fmt.Println("Value for Deposit is: ", readCmdDeposit) //show the user command recived

	return readCmdDeposit

}

package main

import (
	"bank/accounts"
	"bank/holders"
	"fmt"
	"os"
)

func paymentsBills(conta verificarConta, valorDoBoleto float64) {
	conta.Pay(valorDoBoleto)

}

type verificarConta interface {
	Pay(valor float64) string
}

func main() {

	holderAlfredo := holders.Holder{"Alfredo", "129.092.076-17", "Developer", 1800}
	accountAlfredo := accounts.CurrentAccount{holderAlfredo, 001, 4760850, 5, 0260, "Nu Pagamentos S.A - Instituição Financeira", 1800}

	//paymentsBills(&accountAlfredo, 60)
	fmt.Println(accountAlfredo.AccountBalance)

	fmt.Println(accountAlfredo)

	showMenu()

	cdmOp := readCmd()

	switch cdmOp {
	case 1:
		accounts.OpTransfer()
	case 2:
		accounts.OpDeposit()
	case 0:
		fmt.Println("Exit")
		os.Exit(0)
	default:
		fmt.Println("Unrecognized command.")

	}

}

//func for show the menu
func showMenu() {

	fmt.Println(" What operation do you want accomplish?")
	fmt.Println("1- Transfer - TED")
	fmt.Println("2- Deposit")
	fmt.Println("0- Sair do Programa")

}

//func for read command of menu
func readCmd() int {

	var readedCmd int //var for read user command

	fmt.Scan(&readedCmd)                               //recive user command
	fmt.Println("The chosen command was: ", readedCmd) //show the user command recived

	return readedCmd

}

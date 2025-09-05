package main

import (
	"fmt"
	"strconv"
)

func main(){
	deposit, withdraw, getBalance, getInfo := createBankAccount("John Doe", 1000.0)
	fmt.Println(getInfo())

	fmt.Println(deposit(250.0))

	fmt.Println(withdraw(500.0))
	fmt.Println(withdraw(1000.0))

	fmt.Println(getBalance())

	deposit2, _, getBalance2, getInfo2 := createBankAccount("Jane Doe", 1000.0)
	fmt.Println(getInfo2())
	fmt.Println(deposit2(250.0))
	fmt.Println(getBalance2())

}

func createBankAccount(accountHolder string, initialBalance float64) (func (float64) float64, func (float64) float64, func () float64, func () string) {
	balance := initialBalance // balance is a closure variable, it will be stored in the heap
	deposit := func (amount float64) float64 {
		fmt.Println("Depositing", amount, "to", accountHolder)
		balance += amount
		return balance
	}
	withdraw := func (amount float64) float64 {
		if amount > balance {
			return -1
		}
		balance -= amount
		return balance
	}
	getBalance := func () float64 {
		return balance
	}
	getInfo := func () string {
		return "Account info: " + accountHolder + " with balance of " + strconv.FormatFloat(balance, 'f', -1, 64)
	}
	return deposit, withdraw, getBalance, getInfo
}
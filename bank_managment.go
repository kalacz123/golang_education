package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	Balance float64
	Owner string
}

func main(){
	firstAccount := BankAccount{
		AccountNumber: "1",
		Balance: 1000,
		Owner: "Bohdan Kalachynskyi",
	}
	secondAccount := BankAccount{
		AccountNumber: "2",
		Balance: 500,
		Owner: "Anna Babenko",
	}

	deposit(&firstAccount, 100)
	deposit(&secondAccount, 50)
	withdraw(&firstAccount, 100)
	transfer(&firstAccount, &secondAccount, 50)
	fmt.Println(firstAccount)
	fmt.Println(secondAccount)
	updateOwner(&firstAccount, "Bohdan K")
	fmt.Println(firstAccount)


}

func deposit(account *BankAccount, amount float64) error {
	 if amount < 0 {
		return errors.New("amount is less than 0")
	 }
	 fmt.Println("Depositing", amount, "to", account.AccountNumber)
	 account.Balance += amount
	 return nil
}

func withdraw(account *BankAccount, amount float64) error{
	if amount > 0 && amount <= account.Balance {
		fmt.Println("Withdrawing", amount, "from", account.AccountNumber)
		account.Balance -= amount
		return nil
	}
	return errors.New("amount is greater than balance")
}

func transfer(from *BankAccount, to *BankAccount, amount float64) error{
	if amount > 0 && amount <= from.Balance {
		fmt.Println("Transferring", amount, "from", from.AccountNumber, "to", to.AccountNumber)
		from.Balance -= amount
		to.Balance += amount
		return nil
	}
	return errors.New("amount is greater than balance")
}

func updateOwner(account *BankAccount, newOwner string){
	fmt.Println("Updating owner of", account.AccountNumber, "to", newOwner)
	account.Owner = newOwner
}

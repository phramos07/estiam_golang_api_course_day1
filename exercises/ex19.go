package main

import "fmt"

type Account struct {
	Id      int
	Owner   string
	Balance float64
}

func NewAccount(id int, owner string, balance float64) *Account {
	return &Account{
		Id:      id,
		Owner:   owner,
		Balance: balance,
	}
}

func (a *Account) Deposit(value float64) {
	a.Balance += value
}

func (a *Account) Withdraw(value float64) {
	a.Balance -= value
}

func (a *Account) String() string {
	return fmt.Sprintf("Acc Number: %d\nOwner: %s\nBalance: %.2f\n", a.Id, a.Owner, a.Balance)
}

type AccountManager struct {
	nextId int
	// save accounts on slice or map ?

}

func (am *AccountManager) CreateAccount(owner string, initialBalance float64) {
	acc := NewAccount(am.nextId, owner, initialBalance)

	// save account on the map by ID

	am.nextId++
}

/*
Exercise: Building a Basic Banking System

Create a basic banking system that allows users to create accounts, deposit and withdraw funds, and view account details.

Define an Account struct with fields:
		AccountNumber, Owner, and Balance.
Deposit(amount float64): A method to deposit funds into the account.
Withdraw(amount float64) error: A method to withdraw funds from the account, returning an error if the balance is insufficient.
String() string: A method to return a formatted string representation of the account.

Define an AccountManager struct responsible for managing a collection of accounts.
		Collection of accounts (slice or map?)
CreateAccount(owner string, initialBalance float64) *Account: A method to create a new account and add it to the collection.
GetAccount(accountNumber int) (*Account, error): A method to retrieve an account based on its account number.
ListAccounts(): Display details of all accounts in the collection.
Create the main function:

Use the AccountManager methods to create accounts, deposit and withdraw funds, and display account details.*/

func main() {}

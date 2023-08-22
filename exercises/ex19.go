package main

import (
	"fmt"
)

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
	accounts map[int]*Account
}

func NewAccountManager() *AccountManager {
	return &AccountManager{
		accounts: make(map[int]*Account),
	}
}

func (am *AccountManager) CreateAccount(owner string, initialBalance float64) {
	acc := NewAccount(am.nextId, owner, initialBalance)

	// save account on the map by ID
	am.accounts[am.nextId] = acc

	am.nextId++
}

func (am *AccountManager) GetAccount(accNumber int) (*Account, error) {
	found := am.accounts[accNumber]
	if found == nil {
		return nil, fmt.Errorf("Couldn't find account %d", accNumber)
	}
	return found, nil
}

func (am *AccountManager) ListAccounts() {
	for _, account := range am.accounts {
		fmt.Println(account)
	}
}

func (am *AccountManager) DeleteAccount(accNumber int) {
	delete(am.accounts, accNumber)
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

func main() {
	manager := NewAccountManager()

	manager.CreateAccount("Pedro", 0.0)
	// manager.CreateAccount("Maxime", 10000000.0)

	manager.ListAccounts()

	acc, err := manager.GetAccount(1)
	if err == nil {
		fmt.Printf("Account owner: %s", acc.Owner)
	} else {
		fmt.Println(err)
	}
}

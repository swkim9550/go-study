package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// function 에 struct 를 만드는것이 method
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("Can't withdraw")

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) changeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "s account. \nHas: ", a.balance)
}

//func main() {
//	account := accounts.NewAccount("seongwoo")
//	//account.balance = 10000 // balance Account struct 의 private인 부분이다.
//	account.Deposit(10)
//	fmt.Println(account)
//	fmt.Println(account.Balance())
//	err := account.Withdraw(20)
//
//	if err != nil {
//		log.Println(err)
//	}
//
//	fmt.Println(account.Balance(), account.Owner())
//	fmt.Println(account)
//}

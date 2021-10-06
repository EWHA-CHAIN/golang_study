package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	owner   string // private
	balance int    // private
}

// Declare Customizing Error
var errNoMoney = errors.New("can't withdraw : Insufficient Cash")

// NewAccount creates Account (func)
func NewAccount(owner string) *Account { //	함수로 생성자 역할을 대신함 - 새로운 Object를 생성하여 주소값 반환
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account (method)
func (account *Account) Deposit(amount int) { // 형태 : func {receiver} {method_name}, account *Account : Deposit()을 호출한 객체를 사용하도록 함
	account.balance += amount
}

// GetBalance of your account
func (account Account) GetBalance() int {
	return account.balance
}

// WithDraw x amount from your account
func (account *Account) WithDraw(amount int) error {
	if account.balance < amount {
		return errNoMoney
	}
	account.balance -= amount
	return nil // same with null or none
}

// ChangeOwner of the account
func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

// GetOwner of the account
func (account Account) GetOwner() string {
	return account.owner
}

func (account Account) String() string {
	return fmt.Sprint(account.owner, "'s account")
}

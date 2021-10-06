package accounts

// Account Struct
type Account struct {
	owner   string // private
	balance int    // private
}

// NewAccount creates Account
func NewAccount(owner string) *Account { //	함수로 생성자 역할을 대신함 - 새로운 Object를 생성하여 주소값 반환
	account := Account{owner: owner, balance: 0}
	return &account
}

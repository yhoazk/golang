package funding


type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}
// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct w/o worring
	// about wheter it's in the stack or heap: Go figures that out
	// for us.
	return &Fund {
		balance: initialBalance,
	}

}

// Methods start with a receiver, in this case the fund.
func (f *Fund) Balance() int {
	return f.balance
}


func (f *Fund) Withdraw(amount int){
	f.balance -= amount
}

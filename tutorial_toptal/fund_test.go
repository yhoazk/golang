package funding

import "testing"

func BenchmarkingFund(b *testing.B){
	// Adds as many dolars as we have iterations this run
	fund := NewFund(b.N)

	//
	for i:=0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() != 0 {
		b.Error("Balance wasnt zero:", fund.Balance())
	}
}

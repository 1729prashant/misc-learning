package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testBiller[org](
		t,
		orgBiller{Plan: "pro"},
		org{Admin: user{UserEmail: "jaskier@oxenfurt.com"}, Name: "Oxenfurt"},
		3000,
		"jaskier@oxenfurt.com",
	)
	testBiller[user](
		t,
		userBiller{Plan: "basic"},
		user{UserEmail: "vesemir@kaermorhen.com"},
		50,
		"vesemir@kaermorhen.com",
	)
	testBiller[user](
		t,
		userBiller{Plan: "pro"},
		user{UserEmail: "zoltan@mahakam.com"},
		100,
		"zoltan@mahakam.com",
	)
	testBiller[org](
		t,
		orgBiller{Plan: "basic"},
		org{Admin: user{UserEmail: "vernon@temeria.com"}, Name: "Temeria"},
		2000,
		"vernon@temeria.com",
	)
	testBiller[org](
		t,
		orgBiller{Plan: "pro"},
		org{Admin: user{UserEmail: "fringilla@nilfgaard.com"}, Name: "Nilfgaard"},
		3000,
		"fringilla@nilfgaard.com",
	)
}

func testBiller[C customer](
	t *testing.T,
	b biller[C],
	c C,
	expectedAmount float64,
	expectedEmail string,
) {
	currentBill := b.Charge(c)
	name := b.Name()
	if currentBill.Amount != expectedAmount ||
		currentBill.Customer.GetBillingEmail() != expectedEmail {
		t.Errorf(`Test Failed:
name:            %v
biller Type:     %T,
customer Type:   %T,
customer:        %v
expected amount: %v
expected email:  %v
actual amount:   %v
actual email:    %v
---------------------------------
`,
			name,
			b,
			c,
			c,
			expectedAmount,
			expectedEmail,
			currentBill.Amount,
			currentBill.Customer.GetBillingEmail(),
		)
	} else {
		fmt.Printf(`Test Passed:
name:            %v
biller Type:     %T,
customer Type:   %T,
customer:        %v
expected amount: %v
expected email:  %v
actual amount:   %v
actual email:    %v
---------------------------------
`,
			name,
			b,
			c,
			c,
			expectedAmount,
			expectedEmail,
			currentBill.Amount,
			currentBill.Customer.GetBillingEmail(),
		)
	}
}

var withSubmit = true

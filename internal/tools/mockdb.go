package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"scott": {
		AuthToken: "apodScott",
		Username:  "scott",
	},
	"oma": {
		AuthToken: "12abc34",
		Username:  "oma",
	},
}

var mockBankDetails = map[string]BankDetails{
	"scott": {
		Cash:     1000,
		Username: "scott",
	},
	"oma": {
		Cash:     20000,
		Username: "oma",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserBankDetails(username string) *BankDetails {
	time.Sleep(time.Second * 1)

	var clientData = BankDetails{}
	clientData, ok := mockBankDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

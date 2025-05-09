package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"hyper": {
		AuthToken: "hyper147",
		Username: "Hyperion",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"hyper": {
		Coins: 6969,
		Username: "Hyperion",
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

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
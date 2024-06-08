package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]loginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coin:     100,
		Username: "alex",
	},
	"jason": {
		Coin:     300,
		Username: "jason",
	},
	"marie": {
		Coin:     200,
		Username: "marie",
	},
}

func (c *mockDB) GetUserLoginDetails(username string) *loginDetails {

	time.Sleep(time.Second * 1)

	var clientData = loginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (c *mockDB) GetUserCoins(username string) *CoinDetails {

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

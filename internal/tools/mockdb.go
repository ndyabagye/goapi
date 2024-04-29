package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"henry": {
		AuthToken: "123ABC",
		Username:  "henry",
	},
	"alex": {
		AuthToken: "456DEF",
		Username:  "alex",
	},
	"mary": {
		AuthToken: "789GHI",
		Username:  "mary",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"henry": {
		Coins:    100,
		Username: "alex",
	},
	"alex": {
		Coins:    200,
		Username: "alex",
	},
	"mary": {
		Coins:    300,
		Username: "mary",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate DB call
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

package tools

type loginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coin     int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *loginDetails
	GetUserCoins(Username string) *CoinDetails
	SetupDatabase() error
}

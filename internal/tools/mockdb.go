package tools

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
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
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

func (m *mockDB) GetUsers() []string {
	var users []string
	for username := range mockLoginDetails {
		users = append(users, username)
	}
	return users
}

func (m *mockDB) GetUserDetails(username string) *UserDetails {
	var loginData = LoginDetails{}
	loginData, okLogin := mockLoginDetails[username]
	var coinData = CoinDetails{}
	coinData, okCoin := mockCoinDetails[username]

	if !okLogin || !okCoin {
		return nil
	}

	var userDetails = UserDetails{loginData.Username, loginData.AuthToken, coinData.Coins}
	return &userDetails
}

func (m *mockDB) UpdateUserCoins(username string, coins int64) bool {
	var coinData = CoinDetails{}
	coinData, okCoin := mockCoinDetails[username]

	if !okCoin {
		return false
	}

	coinData.Coins = coins
	mockCoinDetails[username] = coinData
	return true
}

func (m *mockDB) NewUser(username string, token string, coins int64) bool {
	var loginData = LoginDetails{token, username}
	var coinData = CoinDetails{coins, username}

	_, okLogin := mockLoginDetails[username]
	_, okCoin := mockCoinDetails[username]

	if okLogin || okCoin {
		return false
	}

	mockLoginDetails[username] = loginData
	mockCoinDetails[username] = coinData
	return true
}

func (m *mockDB) RemoveUser(username string) bool {
	_, okLogin := mockLoginDetails[username]
	_, okCoin := mockCoinDetails[username]

	if !okLogin || !okCoin {
		return false
	}

	delete(mockLoginDetails, username)
	delete(mockCoinDetails, username)
	return true
}

func (m *mockDB) SetupDatabase() error {
	return nil
}

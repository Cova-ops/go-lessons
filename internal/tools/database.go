package tools

type LoginDetails struct {
	AuthToken string `json:"authToken"`
	Username  string `json:"username"`
}

type CoinDetails struct {
	Coins    int64  `json:"coins"`
	Username string `json:"username"`
}

type UserDetails struct {
	Username  string `json:"username"`
	AuthToken string `json:"authToken"`
	Coins     int64  `json:"coins"`
}

type DatabaseInterface interface {
	GetUserDetails(username string) *UserDetails
	GetUsers() []string
	UpdateUserCoins(username string, coins int64) bool
	NewUser(username string, token string, coins int64) bool
	RemoveUser(username string) bool
	SetupDatabase() error
}

func NewDatabase() *DatabaseInterface {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		panic(err)
	}

	return &database
}

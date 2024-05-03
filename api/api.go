package api

// Coin balance Params in the url endpoint
type CoinBalanceParams struct {
	Username string
}

// This is the coin balance response struct/"object" that defines the structure of the data we will send back as a response
type CoinBalanceResponse struct {
	Code    int   // this is the response status code that would sent back
	Balance int64 // this is the balance of the user
}

// This is the error struct/"object" that will be sent back as a response if there is an error when a client tries to access the end point
type Error struct {
	Code    int    // this is the response status code that would sent back
	Message string // this is the error message that would sent back
}

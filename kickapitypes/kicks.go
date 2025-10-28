package kickapitypes

type Kicks struct {
	Data    KicksLeaderboard `json:"data"`
	Message string           `json:"message"`
}

type KicksLeaderboard struct {
	Lifetime []KicksLeaderBoardEntry `json:"lifetime"`
	Month    []KicksLeaderBoardEntry `json:"month"`
	Week     []KicksLeaderBoardEntry `json:"week"`
}

type KicksLeaderBoardEntry struct {
	GiftedAmount int    `json:"gifted_amount"`
	Rank         int    `json:"rank"`
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
}

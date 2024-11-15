package telegram

type UpdatesResponse struct {
	Result []Update `json:"result"`
	Ok     bool     `json:"ok"`
}

type Update struct {
	Message string `json:"message"`
	ID      int    `json:"update_id"`
}

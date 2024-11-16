package telegram

type UpdatesResponse struct {
	Result []Update `json:"result"`
	Ok     bool     `json:"ok"`
}

type Update struct {
	Message *IncomingMessage `json:"message"`
	ID      int              `json:"update_id"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

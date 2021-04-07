package models

type DataRetail struct {
	OrderNumber string `json:"orderNumber"`
	FirstName   string `json:"firstname"`
	PhoneNumber string `json:"phonenumber"`
	Transport   string `json:"transport"`
}

type SendMessageWazzupWhatsApp struct {
	ChannelID string `json:"channelId"`
	ChatType  string `json:"chatType"`
	ChatID    string `json:"chatId"`
	Text      string `json:"text"`
}

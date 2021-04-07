package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GodCratos/whatsApp/models"
)

const (
	apiKeyWazzap = "8494ee682fb44c67b10163b67f2ad3c8"
	urlWazzup    = "https://api.wazzup24.com/v2"
)

func GetActiveChannelWhatsApp() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlWazzup+"/channels", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Basic 8494ee682fb44c67b10163b67f2ad3c8")

	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var bodyJson []map[string]interface{}
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range bodyJson {
		if line["transport"].(string) == "whatsapp" && line["state"].(string) == "active" {
			return line["channelId"].(string)
		}
	}
	return ""
}

func SendMessageInWhatsApp(channelId string, dataRetail models.DataRetail) {
	var sendMessage models.SendMessageWazzupWhatsApp
	sendMessage.ChannelID = channelId
	sendMessage.ChatType = "whatsapp"
	sendMessage.ChatID = dataRetail.PhoneNumber
	sendMessage.Text = fmt.Sprintf("Hello, %v", dataRetail.FirstName)
	jsonByte, err := json.Marshal(sendMessage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonByte))
	a := bytes.NewReader(jsonByte)
	client := http.Client{}
	req, err := http.NewRequest("POST", urlWazzup+"/send_message", a)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Basic 8494ee682fb44c67b10163b67f2ad3c8")
	req.Header.Set("Content-Type", "application/json")

	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GodCratos/whatsApp/models"
	"github.com/GodCratos/whatsApp/services"
)

func GetMessageInWhatsAppHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var dataRetail models.DataRetail
	err = json.Unmarshal(body, &dataRetail)
	if err != nil {
		fmt.Println(err)
	}
	channelId := services.GetActiveChannelWhatsApp()
	services.SendMessageInWhatsApp(channelId, dataRetail)
	var chatWhatsApp models.DataRetail
	chatWhatsApp.FirstName = "Sanya"
	chatJSON, _ := json.Marshal(chatWhatsApp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(chatJSON)
}

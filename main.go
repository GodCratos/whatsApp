package main

import (
	"fmt"
	"net/http"

	"github.com/GodCratos/whatsApp/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get/message", handlers.GetMessageInWhatsAppHandler).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psenna/go-mail/model"
)

func CreateMail(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var mail model.Mail
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(mail)

	// Validate the mail object
	if mail.Subject == "" || mail.Body == "" || mail.From == "" || mail.To == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

}

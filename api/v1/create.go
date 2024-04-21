package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kallepan/go-router-test/internal/model"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// get the request body
	var body model.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new user
	user := model.User{
		Name: body.Name,
	}

	// Write a response to the client
	w.Write([]byte(fmt.Sprintf("User %s created.\n", user.Name)))
}

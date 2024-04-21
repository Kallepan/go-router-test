package api

import (
	"fmt"
	"net/http"

	"github.com/kallepan/go-router-test/internal/model"
)

// GoodbyeHandler is a simple HTTP handler that writes a response.
func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new user
	user := model.User{
		Name: "John Doe",
	}

	// Write a response to the client
	w.Write([]byte(fmt.Sprintf("Goodbye, %s!\n", user.Name)))
}

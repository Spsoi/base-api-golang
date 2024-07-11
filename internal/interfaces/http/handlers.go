package http

import (
	"Azgart/internal/application"
	"net/http"
	"strings"
)

func RegisterHandlers(router *http.ServeMux) {
	userService := &application.UserService{}
	router.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		userID := strings.TrimPrefix(r.URL.Path, "/users/")
		user := userService.GetUserByID(userID)
		w.Write([]byte("User ID: " + user.ID + ", Name: " + user.Name))
	})
}

package server

import (
	"fmt"
	"net/http"
	"reader/internal/configs"
	"reader/internal/usecases"
)

func handleSendUsers(w http.ResponseWriter, r *http.Request) {
	err := usecases.SendUsers()
	if err != nil {
		fmt.Fprintf(w, "error to send users: %s", err)
	}
}

func UP() {
	port := configs.GetAPI().Port
	fmt.Printf("running on port %s\n", port)

	http.HandleFunc("/users", handleSendUsers)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}

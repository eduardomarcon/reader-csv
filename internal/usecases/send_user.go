package usecases

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SendUsers() error {
	chanUser := make(chan User)

	file, err := readFile(chanUser)
	if err != nil {
		return err
	}
	defer file.Close()

	for user := range chanUser {
		userJson, _ := json.Marshal(user)
		fmt.Printf("user %s sended\n", userJson)
	}

	return nil
}

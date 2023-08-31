package login

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"websocket_to_telegram/internal/models"
)

func Auth(server string, credentials string) models.User {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://"+server+".artux.net/pdanetwork/api/v1/user/info", nil)
	if err != nil {
		log.Println("Ошибка при создании GET-запроса:", err)
		return models.User{}
	}

	req.Header.Set("Authorization", credentials)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Ошибка при выполнении GET-запроса:", err)
		return models.User{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении тела ответа:", err)
		return models.User{}
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Ошибка при десериализации данных:", err)
		return models.User{}
	}

	return user
}

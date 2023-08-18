package chat

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetChats(header string) []string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://dev.artux.net/pdanetwork/api/v1/admin/chats/types", nil)
	if err != nil {
		log.Println("Ошибка при создании GET-запроса:", err)
		return []string{}
	}

	req.Header.Set("Authorization", header)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Ошибка при выполнении GET-запроса:", err)
		return []string{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении тела ответа:", err)
		return []string{}
	}

	var chats []string
	err = json.Unmarshal(body, &chats)
	if err != nil {
		log.Println("Ошибка при десериализации данных:", err)
		return []string{}
	}

	return chats
}

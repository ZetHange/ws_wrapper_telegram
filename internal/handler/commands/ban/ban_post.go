package ban

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func BanAlways(header string, uuid string) bool {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://dev.artux.net/pdanetwork/api/v1/admin/bans/"+uuid+"/set/always", nil)
	if err != nil {
		log.Println("Ошибка при создании GET-запроса:", err)
		return false
	}
	req.Header.Set("Authorization", header)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Ошибка при выполнении GET-запроса:", err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении тела ответа:", err)
		return false
	}

	var success bool
	err = json.Unmarshal(body, &success)
	if err != nil {
		log.Println("Ошибка при десериализации данных:", err)
		return false
	}
	return success
}
func BanTime(header string, uuid string, secs int, reason string, message string) bool {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://dev.artux.net/pdanetwork/api/v1/admin/bans/"+uuid+"?secs="+fmt.Sprintf("%v", secs)+"&reason="+reason+"&message="+message, nil)
	if err != nil {
		log.Println("Ошибка при создании GET-запроса:", err)
		return false
	}
	req.Header.Set("Authorization", header)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Ошибка при выполнении GET-запроса:", err)
		return false
	}
	defer resp.Body.Close()
	return true
}

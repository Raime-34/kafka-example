package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getRandomUser() (User, error) {
	var result Result

	resp, err := http.Get("https://randomuser.me/api/")
	if err != nil || resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("ошибка при получении данных")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, fmt.Errorf("ошибка чтения")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return User{}, fmt.Errorf("ошибка парсинга")
	}

	return result.Users[0], nil
}

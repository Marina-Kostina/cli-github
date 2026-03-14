package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	CreatedAt   string `json:"created_at"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Ошибка! Укажите репозиторий в формате: 'владелец/репозиторий'")
		return
	}

	all_string := os.Args[1]

	parts_of_string := strings.Split(all_string, "/")

	if len(parts_of_string) != 2 {
		fmt.Println("Ошибка! Укажите репозиторий в формате: 'владелец/репозиторий'")
		return
	}

	owner := parts_of_string[0]
	repo := parts_of_string[1]

	fmt.Printf("Указанный репозиторий: %s/%s\n", owner, repo)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	fmt.Printf("Обращение по ссылке: %s\n", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка HTTP: %s\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа: %v\n", err)
		return
	}

	repository := Repository{}
	err = json.Unmarshal(body, &repository)
	if err != nil {
		fmt.Printf("Ошибка при парсинге JSON: %v\n", err)
		return
	}

	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("		Полученная информация о репозитории %s:\n", os.Args[1])
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("Имя репозитория: %s\n", repository.Name)
	fmt.Printf("Описание: %s\n", repository.Description)
	fmt.Printf("Количество звёзд: %d\n", repository.Stars)
	fmt.Printf("Количество форков: %d\n", repository.Forks)
	fmt.Printf("Дата создания: %s\n", repository.CreatedAt[:10])
	fmt.Println(strings.Repeat("-", 80))
}

package main

import (
	"fmt"
	"os"
)

func main() {
	token := os.Getenv("TELEGRAM_APITOKEN")
	if token == "" {
		fmt.Println("Переменная окружения TELEGRAM_APITOKEN не установлена")
	} else {
		fmt.Println("Переменная окружения TELEGRAM_APITOKEN установлена:", token)
	}
}

package main

import (
	"io/ioutil"
	"net/http"
)

// Точка входа
func main() {
	botToken := "6211018847:AAEf6Gv_9_1evd2z-4nBlPsbkx2WTnZGQRQ"
	botApi := "https://api.telegram.org/bot"
	//https://api.telegram.org/bot<token>/METHOD_NAME
	botUrl := botApi + botToken
}

func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
}
func respons() {

}

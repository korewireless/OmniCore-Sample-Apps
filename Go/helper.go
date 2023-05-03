package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GenerateToken(clientId string, clientSecret string) string {
	var token = ""

	url := "https://api.korewireless.com/Api/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials&client_id=" + clientId + "&client_secret=" + clientSecret)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return token
	}
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return token
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return token
	}
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	if res.StatusCode != 200 {
		return token
	}
	token = data["access_token"].(string)

	return token

}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(clientId string, clientSecret string) (string, error) {
	var token = ""

	url := "https://api.korewireless.com/Api/token"
	method := "POST"

	if strings.TrimSpace(clientId) == "" || strings.TrimSpace(clientSecret) == "" {
		return "", errors.New("client ID or Secret is invalid")
	}
	payload := strings.NewReader("grant_type=client_credentials&client_id=" + clientId + "&client_secret=" + clientSecret)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return token, err
	}
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return token, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return token, err
	}
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	if res.StatusCode != 200 {
		return token, err
	}
	token = data["access_token"].(string)

	return token, err

}

func IsTokenExpired(accessToken string) bool {
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return true
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println(err)
		return true
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		fmt.Println(err)
		return true
	}

	expTime := time.Unix(int64(exp), 0)
	return time.Now().UTC().After(expTime)
}

func FetchToken(currentToken string, clientID string, clientSecret string) (string, error) {
	if currentToken != "" {
		if IsTokenExpired(currentToken) {
			return GenerateToken(clientID, clientSecret)
		} else {
			return currentToken, nil
		}
	} else {
		return GenerateToken(clientID, clientSecret)
	}
}

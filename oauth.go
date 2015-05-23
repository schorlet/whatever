package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func GetToken(key, secret string) (token string, err error) {
	body, err := fetchToken(key, secret)
	if err != nil {
		return
	}

	token, err = parseToken(body)
	return
}

func parseToken(body []byte) (token string, err error) {
	var auth struct {
		Type  string `json:"token_type"`
		Value string `json:"access_token"`
	}

	err = json.Unmarshal(body, &auth)
	if err != nil {
		return
	}

	if auth.Type != "bearer" {
		err = fmt.Errorf("expected bearer but was %s", auth.Type)
		return
	}

	token = auth.Value
	log.Println("oauth:", token)

	return
}

func parseError(body []byte) (message string, err error) {
	var errors struct {
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	err = json.Unmarshal(body, &errors)
	if err != nil {
		return
	}

	if len(errors.Errors) > 0 {
		message = errors.Errors[0].Message
	}

	return
}

func fetchToken(key, secret string) (body []byte, err error) {
	oauthUrl := BaseUrl + "/oauth2/token"
	form := bytes.NewBuffer([]byte("grant_type=client_credentials"))

	req, err := http.NewRequest("POST", oauthUrl, form)
	if err != nil {
		return
	}

	key = url.QueryEscape(key)
	secret = url.QueryEscape(secret)
	req.SetBasicAuth(key, secret)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Content-Length", "29")

	res, body, err := fetch(req)

	if res.StatusCode != http.StatusOK {
		message, errp := parseError(body)
		if errp != nil {
			return
		}
		err = fmt.Errorf("oauth: %s", message)
	}

	return
}

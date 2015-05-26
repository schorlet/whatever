package main

import (
	"log"
	"net/http"

	. "./backend"
)

func main() {
	key, secret := Credentials()
	token, err := GetToken(key, secret)
	if err != nil {
		log.Fatal(err)
	}

	handler := NewAppHandler(token)
	err = http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := flag.Bool("client", false, "start client")
	server := flag.Bool("server", false, "start server")
	flag.Parse()

	if flag.NFlag() != 1 {
		fmt.Println("go run main.go [-server | -client args]")
		os.Exit(1)
	}

	key, secret := credentials()
	token, err := GetToken(key, secret)
	if err != nil {
		log.Fatal(err)
	}

	if *client {
		runClient(token)
	} else if *server {
		runServer(token)
	}
}

func runClient(token string) {
	args := clientArgs()

	query := strings.Join(args, " ")
	tweets, err := Search(query, token)
	if err != nil {
		log.Fatal(err)
	}

	body, err := json.Marshal(tweets)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, body, "", "\t")
	out.WriteTo(os.Stdout)
}

func runServer(token string) {
	handler := NewAppHandler(token)

	var err = http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal(err)
	}
}

func clientArgs() (args []string) {
	args = flag.Args()

	if len(args) == 0 {
		log.Fatal("main: missing search arguments")
	}

	for i, arg := range args {
		if strings.Index(arg, " ") != -1 {
			args[i] = fmt.Sprintf(`"%s"`, arg)
		}
	}
	return
}

func credentials() (key, secret string) {
	key = os.Getenv("CONSUMER_KEY")
	secret = os.Getenv("CONSUMER_SECRET")

	if key == "" || secret == "" {
		log.Fatal("main: missing consumer credentials")
	}
	return
}

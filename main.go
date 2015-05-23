package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := arguments()
	key, secret := credentials()

	token, err := GetToken(key, secret)
	if err != nil {
		log.Fatal(err)
	}

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

func arguments() (args []string) {
	flag.Parse()
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

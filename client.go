package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	. "./backend"
)

func main() {
	args := clientArgs()

	key, secret := Credentials()
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

func clientArgs() (args []string) {
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

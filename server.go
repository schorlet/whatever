package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    key, secret := credentials()
    token, err := GetToken(key, secret)
    if err != nil {
        log.Fatal(err)
    }

    http.Handle("/search", SearchHandler(token))
    http.Handle("/", http.FileServer(http.Dir("./")))

    err = http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func SearchHandler(token string) http.Handler {
    var fn = func(w http.ResponseWriter, r *http.Request) {
        query := r.FormValue("q")
        if query == "" {
            http.Error(w, "missing q=arguments", http.StatusBadRequest)
            return
        }

        body, err := fetchQuery(query, token)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        w.Write(body)
    }

    return http.HandlerFunc(fn)
}

func credentials() (key, secret string) {
    key = os.Getenv("CONSUMER_KEY")
    secret = os.Getenv("CONSUMER_SECRET")

    if key == "" || secret == "" {
        log.Fatal("server: missing consumer credentials")
    }
    return
}

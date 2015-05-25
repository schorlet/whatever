package backend

import "net/http"

func NewAppHandler(token string) http.Handler {
	router := http.NewServeMux()
	router.Handle("/search", SearchHandler(token))
	router.Handle("/", http.FileServer(http.Dir("./")))
	return router
}

func SearchHandler(token string) http.Handler {
	var fn = func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("q")
		if query == "" {
			http.Error(w, "missing q=arguments", http.StatusBadRequest)
			return
		}

		println(query)
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

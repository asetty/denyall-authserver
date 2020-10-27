package main

import (
	"flag"
	"log"
	"net/http"
)

// TODO configure 403 template path instead of just writing text
func denyHandler(w http.ResponseWriter, r *http.Request) {
	// XXX could log more fields e.g. specific headers,
	log.Printf("received request: method=%q host=%q url=%q ", r.Method, r.Host, r.URL.String())
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Not allowed."))
}

func main() {
	addr := flag.String("addr", ":8080", "address for server to listen on")
	flag.Parse()

	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/", denyHandler)

	log.Printf("Starting deny all auth server at %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

// healthHandler responds to /healthz endpoint for application monitoring
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// expect only GET requests
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unexpected http method"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

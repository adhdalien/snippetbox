package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Make sure we only handle /, and otherwise use http.NotFound() to send
	// a 404 response to the client.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// NOTE: Go's servemux treats the URL pattern "/" like a catch-all.
	// NOTE: Longer URL patterns always take precedence over shorter ones.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// NOTE: The TCP network address here should be in the format "host:port". Right
	// now it's listening on every network interface and port 4000.
	log.Println("Starting Snippetbox server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

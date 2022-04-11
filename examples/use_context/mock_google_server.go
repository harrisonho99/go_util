package use_context

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	MOCK_SERVER_PORT = ":3000"
	MOOK_SERVER_URL  = "localhost"
	SLOW             = 1
)

var data = `{
    "ResponseData": {
        "Results": [
            {
                "TitleNoFormatting": "Concurrency is not parallelism",
                "URL": "https://go.dev/blog/waza-talk"
            },
            {
                "TitleNoFormatting": "Advanced Go Concurrency Patterns",
                "URL": "https://go.dev/blog/io2013-talk-concurrency"
            },
            {
                "TitleNoFormatting": "Program your next server in Go",
                "URL": "https://www.youtube.com/watch?v=5bYO60-qYOI"
            }
        ]
    }
}`

func search(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(time.Second * SLOW)
	fmt.Fprintf(w, data)
}

type GOOGLEServer struct{}

func (s *GOOGLEServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("google_server", r.URL.Path)
	switch r.URL.Path {
	case "/search":
		search(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func UseMockServer() {
	log.Println("UseMockServer")
	server := new(GOOGLEServer)
	log.Fatal(http.ListenAndServe(MOCK_SERVER_PORT, server))
}

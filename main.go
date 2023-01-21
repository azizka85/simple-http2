package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.Int("port", 3000, "The server port")
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Protocol is %s", r.Proto)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	count := 1

	for count < 6 {
		msg := fmt.Sprintf("Hello %d times!\n", count)
		log.Print(msg)
		fmt.Fprint(w, msg)
		time.Sleep(time.Second)

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		count++
	}
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", Info)
	mux.HandleFunc("/hello", Hello)

	fs := http.FileServer(http.Dir("./assets"))

	mux.Handle(
		"/assets/",
		http.StripPrefix("/assets/", fs),
	)

	addr := fmt.Sprintf("localhost:%d", *port)

	log.Printf("Listening %s", addr)

	log.Fatal(
		http.ListenAndServeTLS(
			addr,
			"server.crt",
			"server.key",
			mux,
		),
	)
}

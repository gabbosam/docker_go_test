package main

import (
	"fmt"
	"log"
	"net/http"
)

func httpsWorker(done chan bool) {
	fmt.Print("working...")
	err := http.ListenAndServeTLS(":443", "/etc/server/server.crt", "/etc/server/server.key", nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("https done")

	// Send a value to notify that we're done.
	done <- true
}

func httpWorker(done chan bool) {
	fmt.Print("http working...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "0")
	})

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool)
	go httpWorker(done)
	go httpsWorker(done)
	// Block until we receive a notification from the
	// worker on the channel.
	for i := 0; i < 2; i++ {
		<-done
	}

}

package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	setupApi()

	log.Fatal(http.ListenAndServe(":3001", nil))
	// need to generate server.cert and server.key from openssl
	//log.Fatal(http.ListenAndServeTLS(":3001", "server.cert", "server.key"))
}

func setupApi() {

	ctx := context.Background()

	manager := NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWs)
	http.HandleFunc("/login", manager.loginHandler)
}

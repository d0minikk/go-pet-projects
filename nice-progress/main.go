package main

import (
	"fmt"
	"log"
	"nice-progress/phrases"
	"nice-progress/server"
)

func main() {
	pm := phrases.NewPhraseManager()
	srv := server.NewServer(pm)
	port := "8080"
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(srv.Start(port))
}

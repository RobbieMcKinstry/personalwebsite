package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/RobbieMcKinstry/personal-website/handlers"
)

const PORT = 8080

func main() {

	http.Handle("/static/", handlers.Static())
	http.HandleFunc("/research", handlers.ResearchHandler)
	http.HandleFunc("/engineering", handlers.EngineeringHandler)
	http.HandleFunc("/teaching", handlers.TeachingHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	http.HandleFunc("/other", handlers.OtherHandler)
	http.HandleFunc("/memes", handlers.MemeHandler)
	http.HandleFunc("/keybase.txt", handlers.KeybaseVerificationHandler)

	http.HandleFunc("/", handlers.IndexHandler)

	launchServer()
	wait()
}

func launchServer() {
	fmt.Printf("Server on port %d\n", PORT)
	port := fmt.Sprintf(":%d", PORT)
	go func() {
		log.Fatal(http.ListenAndServe(port, nil))
	}()
}

func wait() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	s := <-interrupt
	fmt.Println("Caught interrupt signal:", s)
}

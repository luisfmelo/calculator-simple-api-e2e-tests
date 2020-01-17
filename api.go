package main

import (
	"github.com/gorilla/mux"
	"github.com/luisfmelo/calculator-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	shutChan := make(chan os.Signal, 1)
	signal.Notify(shutChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.Health)
	router.HandleFunc("/calc", handlers.CalculatorHandler).Methods("POST")

	log.Print("app starting")
	go func() {
		if err := http.ListenAndServe(":9999", router); err != nil {
			log.Fatal(err)
		}
	}()

	<-shutChan

	log.Print("app shutdown")
}

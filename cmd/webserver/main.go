package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ablqk/santorini/api/newgame"
	"github.com/ablqk/santorini/api/gamestate"
	"github.com/ablqk/santorini/api"
	"github.com/ablqk/santorini/api/play"
)

func main() {

	port := flag.Int("port", 3000, "the port to start the web app on")
	flag.Parse()

	server := server{
		router: mux.NewRouter().StrictSlash(true),
	}

	server.handle(newgame.NewEndpoint())
	server.handle(gamestate.NewEndpoint())
	server.handle(play.NewEndpoint())

	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), server.router))

}

type server struct {
	router *mux.Router
}

func (s server) handle(endpoint api.Endpoint) {
	s.router.Handle(endpoint.Path(), api.NewHandler(endpoint)).Methods(endpoint.Verb())
}

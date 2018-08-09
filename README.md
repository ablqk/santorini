A RESTful-API-based version of the board game Santorini, by Dr. Gordon Hamilton, created without any sort of permission.

# Quickstart

	go run cmd/webserver/main.go

# Play in a terminal

	go run cmd/cli/main.go

# Work in progress : TO DO

## Server side

### API documentation
* Update yaml file
* Create quickstart guide

### General
* Unit tests please
* Fix `http: multiple response.WriteHeader calls`

### Endpoints

### Play one turn
* move a pawn

## CLI client
The idea is to implement a client to play in a console. Because why not.

### Client
* Manage errors
* Connect as second player
* Play as second player
* Refresh as first player
* Check end of game
* Build dynamic IP:PORT url for HTTP calls


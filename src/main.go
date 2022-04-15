package main

import (
	"./service"
	"github.com/urfave/negroni"
)

func main() {
	router := service.InitializeRouter()
	server := negroni.Classic()
	server.UseHandler(router)

	port := "localhost:8080"
	server.Run(port)
}

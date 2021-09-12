package main

import (
	"flag"
	"fmt"
	"net/http"
	"paysolut/handler"
)

func main() {
	var (
		err error

		port *int

		server handler.ServerType
	)

	port = flag.Int("p", 9080, "service port")
	flag.Parse()

	fmt.Println("service run on port", *port)
	fmt.Println("to stop the service, press [Ctrl+C]")

	http.HandleFunc("/", server.Handler)

	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Println("error:", err)
	}
}

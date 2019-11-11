package main

import (
	"os"

	"github.com/whichxjy/Service-Computing-Homework/cloudgo/service"
	flag "github.com/spf13/pflag"
)

const defaultPort string = "8888"

func main() {
	// try to retrieve the value of the "PORT" environment variable
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	// if the use set the port, then use what user set
	pPort := flag.StringP("port", "p", "", "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// run server
	server := service.NewServer()
	server.Run(":" + port)
}

package main

import (
	"authentication/models"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	port = flag.String("port", "8080", "port to be used")
	ip   = flag.String("ip", "localhost", "ip to be used")
)

func main() {
	flag.Parse()
	flags := models.NewFlags(*ip, *port)
	fmt.Println("Hello, world.")

	logger := log.New(os.Stdout, "auth", 1)
	route := route.Regi
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	portPtr := flag.Int("port", 3001, "Port to run application on")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *portPtr)

	log.Printf("Listening on %s", addr)

	app := NewApp(os.Getenv("DATABASE_URL"))
	app.Run(addr)
}

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	portPtr := flag.Int("port", 3001, "Port to run application on")
	addr := fmt.Sprintf(":%d", *portPtr)

	app := NewApp(os.Getenv("DATABASE_URL"))
	app.Run(addr)
}

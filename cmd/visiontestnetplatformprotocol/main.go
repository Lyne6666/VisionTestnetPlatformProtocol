// cmd/visiontestnetplatformprotocol/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"visiontestnetplatformprotocol/internal/visiontestnetplatformprotocol"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := visiontestnetplatformprotocol.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

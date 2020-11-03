package main

import (
	"flag"
	"log"
	mysql_scanner "mysql-scanner"
)

const DefaultAddress = "127.0.0.1"
const DefaultPort = 3306

func main() {
	var addr = flag.String("h", DefaultAddress, "network address to connect to")
	var port = flag.Int("p", DefaultPort, "port to connect to")
	flag.Parse()

	client := mysql_scanner.Client{}
	greeting, err := client.Probe(*addr, *port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n %s", greeting)
}

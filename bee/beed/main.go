package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time" // or "runtime"

	"github.com/Emyrk/LendingBot/bee"
)

func main() {
	var (
		address = flag.String("a", "localhost:7000", "Address to connect to the balancer")
		dba     = flag.String("dba", "mongodb://localhost:27017", "Address for db to connect")
		dbu     = flag.String("dbu", "", "Username for db to connect")
		dbp     = flag.String("dbp", "", "Password for db to connect")
	)

	flag.Parse()

	be := bee.NewBee(*address, dba, dbu, dbp)
	be.Run()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("Recievived ctrl+c. Closing balancer")
		be.Shutdown()
		os.Exit(1)
	}()

	log.Printf("Now running Bee")
	for {
		log.Printf("Using %d users", len(be.Users))
		time.Sleep(10 * time.Minute)
	}
}
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/namelew/VirtualWallet/internal/databases"
	"github.com/namelew/VirtualWallet/internal/envoriment"
	"github.com/namelew/VirtualWallet/internal/router"
)

func main() {
	envoriment.LoadFile(".env")
	db := databases.New()

	db.Connect()
	db.Migrate()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		db.Disconnect()
		os.Exit(1)
	}()

	router.Route()
}

package main

import (
	"github.com/namelew/VirtualWallet/internal/databases"
	"github.com/namelew/VirtualWallet/internal/envoriment"
)

func main() {
	envoriment.LoadFile(".env")
	db := databases.New()

	db.Connect()
	db.Migrate()

	//db.ClientTransfer(3, 4, 100)

	db.Disconnect()
}

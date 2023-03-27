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

	// db.AddClient(clients.Client{
	// 	Name:   "Vanderlei",
	// 	Amount: 100,
	// })

	// db.AddClient(clients.Client{
	// 	Name:   "Agnaldo",
	// 	Amount: 110,
	// })

	db.ClientTransfer(3, 4, 110)

	db.Disconnect()
}

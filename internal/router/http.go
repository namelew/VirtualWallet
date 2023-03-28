package router

import (
	"github.com/labstack/echo"
	"github.com/namelew/VirtualWallet/internal/clients"
	"github.com/namelew/VirtualWallet/internal/envoriment"
	"github.com/namelew/VirtualWallet/internal/transations"
)

func Route() {
	router := echo.New()

	router.GET("/:id/amount", clients.GetSaldo)
	router.POST("/transfer/:source/:target/:val", transations.Tranfer)

	router.Logger.Fatal(router.Start(":" + envoriment.GetVar("PORT")))
}

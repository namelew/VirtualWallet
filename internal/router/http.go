package router

import (
	"github.com/labstack/echo"
	"github.com/namelew/VirtualWallet/internal/envoriment"
)

func Route() {
	router := echo.New()

	router.GET("/:id/amount", nil)
	router.POST("/transfer/:source/:target/:val", nil)

	router.Logger.Fatal(router.Start(":" + envoriment.GetVar("PORT")))
}

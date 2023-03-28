package router

import (
	"github.com/labstack/echo"
	"github.com/namelew/VirtualWallet/internal/controllers"
	"github.com/namelew/VirtualWallet/internal/databases"
	"github.com/namelew/VirtualWallet/internal/envoriment"
)

func Route(d *databases.Database) {
	router := echo.New()
	controller := controllers.New(d)

	router.GET("/clients/:id/amount", controller.Get)
	router.POST("/transitions/transfer/:source/:target/:val", controller.Add)

	router.Logger.Fatal(router.Start(":" + envoriment.GetVar("PORT")))
}

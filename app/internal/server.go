package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/sho0126hiro/toriton/app/internal/interface/handler/rest"
	"github.com/sho0126hiro/toriton/app/internal/openapi/v1"
)

func NewServer() (*echo.Echo, *AppServer) {
	e := echo.New()
	appServer := new(AppServer)
	return e, appServer
}

// AppServer ServerInterfaceを実装
type AppServer struct {
	rest.TestHandler
	rest.UserHandler
}

func Run() {
	router, server := NewServer()
	openapi.RegisterHandlers(router, server)
	router.Logger.Fatal(router.Start(":8080"))
}

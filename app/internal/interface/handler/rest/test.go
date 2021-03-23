package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sho0126hiro/toriton/app/internal/usecase/interactor"
)

type TestHandler struct {
}

func NewTestHandler() *UserHandler {
	return &UserHandler{
		Interactor: interactor.UserInteractor{},
	}
}

func (uh *UserHandler) ApiTest(ctx echo.Context) error {
	fmt.Println("Test")
	return ctx.String(http.StatusOK, "ok")
}

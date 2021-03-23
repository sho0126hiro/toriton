package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sho0126hiro/toriton/app/internal/openapi/v1"
	"github.com/sho0126hiro/toriton/app/internal/usecase/interactor"
)

type UserHandler struct {
	Interactor interactor.UserInteractor
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		Interactor: interactor.UserInteractor{},
	}
}

func (uh *UserHandler) CreateUser(ctx echo.Context) error {
	fmt.Println("CreateUser")
	var req openapi.CreateUserJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	res := &openapi.CreateUserResponse{}
	return ctx.JSON(http.StatusOK, res)
}

package http

import (
	"net/http"

	"github.com/labstack/echo"
	entity "github.com/wahyurudiyan/warehouse-stock/entity/stock"
	"github.com/wahyurudiyan/warehouse-stock/usecase/stock"
	"golang.org/x/net/context"
)

type handler struct {
	usecase stock.Usecase
}

type Handler interface {
	CreateStock(c echo.Context) error
}

func NewHandler(u stock.Usecase) Handler {
	return &handler{u}
}

func (h *handler) CreateStock(c echo.Context) error {
	var req entity.Request
	ctx := context.Background()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.usecase.CreateStock(ctx, req.Data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "OK")
}

package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/evzpav/gorsk/pkg/api/trade"

	gorsk "github.com/evzpav/gorsk/pkg/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc trade.Service
}

// NewHTTP creates new trade http service
func NewHTTP(svc trade.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/trades")
	ur.POST("", h.create)
	ur.GET("", h.list)
	ur.GET("/:id", h.view)
	ur.PATCH("/:id", h.update)
	ur.DELETE("/:id", h.delete)
}

// Custom errors
var (
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
)

// User create request

func (h *HTTP) create(c echo.Context) error {
	r := new(gorsk.Trade)

	if err := c.Bind(r); err != nil {
		return err
	}

	trade, err := h.svc.Create(c, *r)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &trade)
}

type listResponse struct {
	Trades []gorsk.Trade `json:"trades"`
	Page   int           `json:"page"`
}

func (h *HTTP) list(c echo.Context) error {
	p := new(gorsk.PaginationReq)
	if err := c.Bind(p); err != nil {
		return err
	}

	result, err := h.svc.List(c, p.Transform())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result, p.Page})
}

func (h *HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return gorsk.ErrBadRequest
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// User update request

func (h *HTTP) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return gorsk.ErrBadRequest
	}

	var t gorsk.Trade
	if err := c.Bind(&t); err != nil {
		return err
	}

	t.ID = id
	updatedTrade, err := h.svc.Update(c, &t)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatedTrade)
}

func (h *HTTP) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return gorsk.ErrBadRequest
	}

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

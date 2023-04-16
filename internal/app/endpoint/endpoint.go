package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}
func (e *Endpoint) Status(c echo.Context) error {

	response := fmt.Sprintf("Days left: %d", e.s.DaysLeft())
	return c.String(http.StatusOK, response)
}

package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Routes() http.Handler {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	group := e.Group("/v1")
	group.POST("/register", s.register)

	return e
}

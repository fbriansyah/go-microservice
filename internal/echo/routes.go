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

	v1 := e.Group("/v1")
	v1.POST("/register", s.register)
	v1.POST("/login", s.login)

	return e
}

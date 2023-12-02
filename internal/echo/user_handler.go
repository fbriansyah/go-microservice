package echo

import (
	"context"
	"net/http"
	"time"

	"github.com/fbriansyah/go-microservice/internal/application/commands"
	"github.com/labstack/echo/v4"
)

func (s *Server) register(c echo.Context) error {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, generateErrorResponse(err.Error()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.app.CreateUser(ctx, commands.CreateUserCmd{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, generateErrorResponse(err.Error()))
	}

	return c.JSON(
		http.StatusOK,
		generateSuccessResponse(
			map[string]any{
				"name":  req.Name,
				"email": req.Email,
			},
		),
	)

}

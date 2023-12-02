package echo

import (
	"context"
	"net/http"
	"time"

	"github.com/fbriansyah/go-microservice/internal/application/commands"
	"github.com/fbriansyah/go-microservice/internal/application/queries"
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
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, generateErrorResponse("name, email and password are required"))
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

func (s *Server) login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, generateErrorResponse(err.Error()))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err := s.app.FindUserByEmail(ctx, queries.FindUserByEmailQuery{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, generateErrorResponse(err.Error()))
	}

	// TODO: Create JWT token

	// TODO: Save JWT token in session manager

	return c.JSON(
		http.StatusOK,
		generateSuccessResponse(
			map[string]any{
				"user": map[string]any{
					"name":  user.Name,
					"email": user.Email,
				},
			},
		),
	)
}

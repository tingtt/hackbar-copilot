package authadapter

import (
	"context"
	"errors"
	"hackbar-copilot/internal/interface-adapter/handler/middleware"
	"slices"
)

type JWTAdapter interface {
	GetEmail(ctx context.Context) (string, error)
	HasBartenderRole(ctx context.Context) bool
}

var ErrUnauthorized = errors.New("unauthorized")

func New() JWTAdapter {
	return &jwtAdapter{}
}

type jwtAdapter struct{}

// GetEmail implements JWTAdapter.
func (j *jwtAdapter) GetEmail(ctx context.Context) (string, error) {
	claims, err := middleware.GetJWT(ctx)
	if err != nil {
		return "", ErrUnauthorized
	}
	return claims.Email, nil
}

// HasBartenderRole implements JWTAdapter.
func (j *jwtAdapter) HasBartenderRole(ctx context.Context) bool {
	claims, err := middleware.GetJWT(ctx)
	if err != nil {
		return false
	}

	return slices.Contains(claims.Roles, "bartender")
}

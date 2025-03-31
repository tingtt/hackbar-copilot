package authadapter

import (
	"context"
	"errors"
	"fmt"
	"hackbar-copilot/internal/interface-adapter/handler/middleware"
	"slices"
)

type JWTAdapter interface {
	GetEmail(ctx context.Context) (string, error)
	GetNameFromOAuth2Provider(ctx context.Context) (string, error)
	HasBartenderRole(ctx context.Context) bool
}

var (
	ErrUnauthorized          = errors.New("unauthorized")
	ErrInvalidJWTClaimFormat = errors.New("invalid JWT claim format")
)

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

// GetNameFromOAuth2Provider implements JWTAdapter.
func (j *jwtAdapter) GetNameFromOAuth2Provider(ctx context.Context) (string, error) {
	claims, err := middleware.GetJWT(ctx)
	if err != nil {
		return "", ErrUnauthorized
	}
	if claims.GitHub != nil {
		return claims.GitHub.ID, nil
	}
	if claims.Google != nil {
		return claims.Google.Username, nil
	}
	return "", fmt.Errorf("%w: claim not contains user info obtained from oauth2 provider", ErrInvalidJWTClaimFormat)
}

// HasBartenderRole implements JWTAdapter.
func (j *jwtAdapter) HasBartenderRole(ctx context.Context) bool {
	claims, err := middleware.GetJWT(ctx)
	if err != nil {
		return false
	}

	return slices.Contains(claims.Roles, "bartender")
}

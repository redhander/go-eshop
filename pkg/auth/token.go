package auth

import (
	"time"

	"github.com/google/uuid"
	"github.com/redhander/go-eshop/internal/db/repository"
)

type TokenGenerator interface {
	GenerateToken(userID uuid.UUID, username string, role repository.UserRole, duration time.Duration) (string, *TokenPayload, error)
	VerifyToken(token string) (*TokenPayload, error)
}

package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("无效Token")
	ErrExpiredToken = errors.New("过期Token")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	RoleKey   []string  `json:"roleKey"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func NewPayload2(username string, roleKey []string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		RoleKey:   roleKey,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

package auth

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

type TokenMaker struct {
	pasetoParser paseto.Parser
	privateKey   paseto.V4AsymmetricSecretKey
	publicKey    paseto.V4AsymmetricPublicKey
}

func NewTokenMaker(privateKeyHex string) (*TokenMaker, error) {
	privateKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("Invalid private key: %v", err)
	}

	return &TokenMaker{
		pasetoParser: paseto.NewParser(),
		privateKey:   privateKey,
		publicKey:    privateKey.Public(),
	}, nil
}

func (m *TokenMaker) CreateToken(userId uuid.UUID, duration time.Duration) (string, error) {
	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(duration))
	token.SetIssuer("xpose")

	token.SetString("user_id", userId.String())

	return token.V4Sign(m.privateKey, nil), nil
}

func (m *TokenMaker) VerifyToken(tokenString string) (uuid.UUID, error) {
	token, err := m.pasetoParser.ParseV4Public(m.publicKey, tokenString, nil)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Invalid token: %v", err)
	}

	idStr, err := token.GetString("user_id")
	if err != nil {
		return uuid.Nil, fmt.Errorf("Invalid. Token is missing 'user_id'")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Invalid 'user_id' format.")
	}

	return id, nil
}

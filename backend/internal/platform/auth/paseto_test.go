package auth

import (
	"testing"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	randomKey := paseto.NewV4AsymmetricSecretKey()
	maker, err := NewTokenMaker(randomKey.ExportHex())
	require.NoError(t, err)

	userId := uuid.New()
	duration := time.Minute

	token, err := maker.CreateToken(userId, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payloadUser, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, userId, payloadUser)
}

func TestExpiredPasetoToken(t *testing.T) {
	randomKey := paseto.NewV4AsymmetricSecretKey()
	maker, err := NewTokenMaker(randomKey.ExportHex())
	require.NoError(t, err)

	userId := uuid.New()

	token, err := maker.CreateToken(userId, -time.Minute)
	require.NoError(t, err)

	_, err = maker.VerifyToken(token)
	require.Error(t, err)

}

package user

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rikkunn23/kokoiko-app-backend-bff/config"
)

type UserClaims struct {
	UserID int64 `json:"user_id"`
	Email string `json:"email"`
	Tel   string `json:"tel"`
	jwt.RegisteredClaims
}

// GenerateJWT ...
func GenerateJWT(email, tel string, userID int64) (string, error) {
	now := time.Now()

	// TODO 電話番号・メールアドレスなどの個人情報は、JWTに含めるべきではないため、必要に応じて削除すること
	claims := UserClaims{
		UserID: userID,
		Email: email,
		Tel:   tel,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)), // 有効期限：24時間
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "kokoiko-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Get().App.JWTSecret))
}

// GenerateRefreshToken ...
func GenerateRefreshToken() (token string, expiresAt time.Time, err error) {

	// TODO時刻はconfigで管理する
	const RefreshTokenDuration = 30 * 24 * time.Hour

	b := make([]byte, 32) // 256bitのランダムバイト
	_, err = rand.Read(b)
	if err != nil {
		return "", time.Time{}, err
	}
	token = base64.URLEncoding.EncodeToString(b)
	expiresAt = time.Now().Add(RefreshTokenDuration)
	return token, expiresAt, nil
}

package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	tokenTTLInMinutes time.Duration
	secretKey         string
}

func NewAuthService(ttl int, secret string) *AuthService {
	return &AuthService{
		tokenTTLInMinutes: time.Duration(ttl) * time.Minute,
		secretKey:         secret,
	}
}

// Claims — структура утверждений, которая включает стандартные утверждения
// и одно пользовательское — UserID
type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

// BuildJWTString создаёт токен и возвращает его в виде строки.
func (as *AuthService) CreateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(as.tokenTTLInMinutes)),
		},
		UserID: userID,
	})

	// создаём строку токена
	tokenString, err := token.SignedString(as.secretKey)
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}

func (as *AuthService) GetUserID(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(as.secretKey), nil
		})
	if err != nil {
		return "", fmt.Errorf("error getting user id from token, err: %s", err.Error())
	}

	if !token.Valid {
		fmt.Println("Token is not valid")
		return "", fmt.Errorf("token is not valid")
	}

	fmt.Println("Token is valid")
	return claims.UserID, nil
}

// struct to pass user id through request context.
type requestContextUserIDValue struct{}

// SetUserID returns context with user id value in returned context.
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, requestContextUserIDValue{}, userID)
}

// GetUserID returns user ID from context.
func (as *AuthService) UserLoginFromCtx(ctx context.Context) (userLogin string) {
	userLogin = ctx.Value(requestContextUserIDValue{}).(string)
	return
}

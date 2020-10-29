package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"ice/global"
	"time"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type JWT struct {
	SigningKey []byte
}

type BaseClaims struct {
	data interface{}
	jwt.StandardClaims
}

func InitClaims(data interface{}) *BaseClaims {
	return &BaseClaims{
		data: data,
	}
}

func (j *JWT) NewJWT() *JWT {
	return &JWT{
		[]byte(global.IceConfig.JWT.Key),
	}
}

func (j *JWT) CreateToken(claims *BaseClaims) (string, error) {
	claims.StandardClaims.NotBefore = time.Now().Unix() - 1000
	claims.StandardClaims.ExpiresAt = time.Now().Unix() + global.IceConfig.JWT.Expire
	claims.StandardClaims.Issuer = global.IceConfig.JWT.Domain
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) VerifyToken(tokenString string) (*BaseClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &BaseClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*BaseClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}

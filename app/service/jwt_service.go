package service

import (
	"github.com/dgrijalva/jwt-go"
	"ice/global"
	"ice/utils/response"
	"time"
)

type JWT struct {
	SigningKey []byte
}

type BaseClaims struct {
	Data interface{}
	BufferTime int64
	jwt.StandardClaims
}

func InitClaims(data interface{}, buffertime int64) *BaseClaims {
	return &BaseClaims{
		Data: data,
		BufferTime: buffertime,
	}
}

func NewJWT() *JWT {
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

func (j *JWT) VerifyToken(tokenString string) (*BaseClaims, int) {
	token, err := jwt.ParseWithClaims(tokenString, &BaseClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, response.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, response.TokenExpire
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, response.TokenBefore
			} else {
				return nil, response.TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*BaseClaims); ok && token.Valid {
			return claims, 0
		}
		return nil, response.TokenInvalid
	} else {
		return nil, response.TokenInvalid
	}

}

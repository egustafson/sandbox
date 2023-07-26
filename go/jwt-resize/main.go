package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtClaim1 struct {
	Name string
	jwt.StandardClaims
}

type JwtClaim2 struct {
	Name  string
	Oname string
	jwt.StandardClaims
}

type JwtKey struct {
	ID        uuid.UUID
	Secret    string
	CreatedAt time.Time
}

var (
	SigningKey = NewJwtKey()
)

func NewJwtKey() (k JwtKey) {
	secret := make([]byte, 32) // 256 bits
	rand.Read(secret)          // inject randomness
	k = JwtKey{
		ID:        uuid.New(),
		Secret:    base64.URLEncoding.EncodeToString(secret),
		CreatedAt: time.Now(),
	}
	return k
}

func getJwtSigningKey(token *jwt.Token) (interface{}, error) {
	return SigningKey.Key(), nil
}

func (k JwtKey) Key() (key []byte) {
	key, err := base64.URLEncoding.DecodeString(k.Secret)
	if err != nil {
		panic(err.Error())
	}
	return key
}

func MakeJwtClaim1(name string) string {
	claims := &JwtClaim1{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(96 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, tokErr := token.SignedString(SigningKey.Key())
	if tokErr != nil {
		panic(tokErr.Error())
	}
	return signedToken
}

func MakeJwtClaim2(name, oname string) string {
	claims := &JwtClaim2{
		Name:  name,
		Oname: oname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(96 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, tokErr := token.SignedString(SigningKey.Key())
	if tokErr != nil {
		panic(tokErr.Error())
	}
	return signedToken
}

func Validate1(signedToken string) *JwtClaim1 {
	token, parseErr := jwt.ParseWithClaims(signedToken, &JwtClaim1{}, getJwtSigningKey)
	if parseErr != nil {
		panic(parseErr.Error())
	}

	if !token.Valid {
		panic("token invalid")
	}

	claims, ok := token.Claims.(*JwtClaim1)
	if !ok {
		panic("failed to cast tokens claims to *JwtClaim1")
	}
	return claims
}

func Validate2(signedToken string) *JwtClaim2 {
	token, parseErr := jwt.ParseWithClaims(signedToken, &JwtClaim2{}, getJwtSigningKey)
	if parseErr != nil {
		panic(parseErr.Error())
	}

	if !token.Valid {
		panic("token invalid")
	}

	claims, ok := token.Claims.(*JwtClaim2)
	if !ok {
		panic("failed to cast tokens claims to *JwtClaim2")
	}
	return claims
}

func main() {
	fmt.Println("run go test")
}

package helper

import (
	"alimed_go/pkg/models"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func loadECDSAPrivateKey() (*ecdsa.PrivateKey, error) {
	keyData, err := os.ReadFile("ecdsa_private.pem")
	if err != nil {
		return nil, fmt.Errorf("could not read private key file: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA private key: %w", err)
	}

	return privateKey, nil
}

type MyClaims struct {
	User models.User `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(values models.User) (string, error) {
	claims := MyClaims{
		User: values,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	// key := os.Getenv("JWT_KEY")
	// println(key)

	byteKey, err := loadECDSAPrivateKey()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(byteKey)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token string) (models.User, error) {
	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")),
			nil
	})
	if err != nil {
		return models.User{}, err
	} else if claims, ok := t.Claims.(*MyClaims); ok {
		//fmt.Println(claims.User, claims.RegisteredClaims.Issuer)
		return claims.User, nil
	} else {
		log.Println("Invalid token")
		return models.User{}, nil
	}
}

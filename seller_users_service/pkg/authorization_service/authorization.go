package service

import (
	"encoding/base64"
	"log"
	"time"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)


type tokenClaims struct {
	jwt.StandardClaims
	Id     string `json:id`
}

func hashPassword(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hashByte), err
}


func (s *Service) GenerateToken(business_id string) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodES512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.expireTime).Unix(),
		},
		business_id,
	})

	key, err := base64.URLEncoding.DecodeString(s.signingKey)
	if err != nil{
		log.Println(err)
		return "", errors.New("Unable to generate token")
	}

	token, err := tkn.SignedString(key)
	if err != nil{
		log.Println(err)
		return "", err
	}

	return token, err
}

func (s *Service) ParseToken(accessToken string) (string, error){
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		key, err := base64.URLEncoding.DecodeString(s.signingKey)
		if err != nil {
			log.Println(err)
		}

		return key, nil
	})

	if err != nil {
		log.Println("Parse token error: ", err)
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.Id, nil
	} else {
		return "", errors.New("Error invalid access token")
	}
}



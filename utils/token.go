package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin/json"
)

type TokenJson struct {
	Exp int64 `json:"exp"`
	Iat int64 `json:"iat"`
	UserId string`json:"userId"`
	RoundKey string `json:"roundKey"`
}

func CreateToken(key string, userId string, roundKey string, expTime int) (string, int64) {
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now();
	nowInt := now.UTC().Unix();
	ext := now.Add(time.Duration(expTime) * time.Millisecond)
	extInt := ext.UTC().Unix()
	content, err := json.Marshal(TokenJson{
	 	extInt,
		nowInt,
		userId,
		strconv.FormatInt(nowInt, 10) + roundKey,
	})
	fmt.Println("json加密：", string(content))
	if err != nil {
		panic(err)
	}
	token.Claims = jwt.MapClaims {
		"content": string(content),
	}
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString, extInt
}

func ParseToken(tokenString string, key string) (interface{}, bool){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

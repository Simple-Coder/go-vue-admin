package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//https://pkg.go.dev/github.com/dgrijalva/jwt-go#example-ParseWithClaims-CustomClaimsType
//go get -u  github.com/dgrijalva/jwt-go
func main() {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		UserName string `json:"userName"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"zhangsan",
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 3,
			Issuer:    "xiedong",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	tokenString := ss

	//type MyCustomClaims struct {
	//	UserName string `json:"userName"`
	//	jwt.StandardClaims
	//}

	// sample token is expired.  override time so it parses as valid
	//at(time.Unix(0, 0), func() {
	//过期测试
	//time.Sleep(5 * time.Second)
	restoken, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := restoken.Claims.(*MyCustomClaims); ok && restoken.Valid {
		fmt.Printf("%v %v", claims.UserName, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	fmt.Println(restoken.Claims.(*MyCustomClaims).UserName)

	//})

}

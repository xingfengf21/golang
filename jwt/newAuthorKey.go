package main

// import (
// 	"fmt"
// 	jwt "github.com/dgrijalva/jwt-go"
// 	"math/rand"
// 	"reflect"
// 	"time"
// )

import (
	"fmt"
	"math/rand"
)

// func newAuthorKey(mySigningKey string, kid string, user string) (string, error) {
// 	// Create the token
// 	n := time.Now().Unix()
// 	rnd := fmt.Sprint(rand.Int31())
// 	fmt.Println("rnd", rnd)
// 	fmt.Println(len(rnd))
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"iss":      "bfupdate",
// 		"username": user,
// 		"iat":      n,
// 		"kid":      kid,
// 		"rnd":      rnd,
// 	})
// 	fmt.Println("token", token)
// 	// header:{}.payload:{iss:user, iat:time, exp:time, rnd:random}.digest

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString(([]byte)(mySigningKey))
// 	return tokenString, err
// }

func main() {
	// token, _ := newAuthorKey("mySigningKey", "kid", "user")
	// fmt.Println("token", token)
	for i := 0; i < 1000; i++ {
		fmt.Println("rand.Int31", rand.Int31())
	}

	// fmt.Println(len(token))
	// fmt.Println(reflect.TypeOf(token))
	// n := time.Now().Unix()
	// fmt.Println("n", n)
}

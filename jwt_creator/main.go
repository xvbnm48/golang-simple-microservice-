package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwgo.io"
	claims[""] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		fmt.Errorf("something when wrong:  %s", err.Error())
		return "", err
	}
	return tokenString, nil

}
func index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)

	if err != nil {
		fmt.Println("failed to generate token")
	}

	fmt.Fprintf(w, string(validToken))
}
func handleRequest() {
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}

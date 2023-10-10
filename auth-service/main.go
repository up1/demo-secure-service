package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

func main() {
	handleRequests()
}

func handleRequests() {
	fmt.Println("Start Identity service...")
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	// For statful => store token in database

	fmt.Fprintf(w, string(validToken))
}

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "from_user"
	claims["aud"] = "user.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %v\n", err)
		return "", err
	}

	return tokenString, nil
}

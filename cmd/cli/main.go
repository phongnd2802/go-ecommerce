package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/phongnd2802/go-ecommerce/pkg/utils"
)

func main() {
	privateKey, publicKey, _ := utils.GenerateRSAKeyPair(2048)
    
    payload := map[string]any {
        "id": "123",
        "email": "2222",
    }
    tokenString, _ := utils.CreateAccessToken(payload, privateKey)
    fmt.Println("Token::", tokenString)

    verifiedToken, err := utils.VerifyToken(tokenString, publicKey)
    if err != nil {
        fmt.Printf("Failed to verify JWT token: %v\n", err)
    } else {
        fmt.Println("JWT Token verified successfully.")
        if claims, ok := verifiedToken.Claims.(jwt.MapClaims); ok && verifiedToken.Valid {
            fmt.Printf("Claims: %v\n", claims)
        }
    }
}

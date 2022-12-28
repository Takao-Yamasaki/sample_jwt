package main

import (
	"context"
	"fmt"
	
	"google.golang.org/api/idtoken"
)

// IDトークンからヘッダー・ペイロード・署名を入手する
func main() {
	googleClientID := "xxxxxxxxxxxxxxxxxx"
	idToken := "xxxxxxxxxxxxxxxxxx"

	tokenValidater, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	// Validateメソッドによる検証作業
	payload, err := tokenValidater.Validate(context.Background(), idToken, googleClientID)
	if err != nil {
		fmt.Println("validate err: ", err)
		return
	}

	fmt.Println(payload.Claims["name"])
}

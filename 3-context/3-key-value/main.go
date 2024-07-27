// o contexto pode passar dados de um lado para outro
package main

import (
	"context"
	"fmt"
)

type contextKey string

const tokenKey contextKey = "token"

func main() {
	ctx := context.Background()
	// em qualquer lugar onde o token for usado, essa informação pode ser recuperada
	ctx = context.WithValue(ctx, tokenKey, "Bearer 123456")
	bookHotel(ctx)
}

// por convenção, o valor do contexto é passado como primeiro argumento
func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println("token:", token)
}

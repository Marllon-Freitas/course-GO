// simular um sistema de reserva que com um tempo limite
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()                            // inicia um contexto em branco
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // cria um contexto com tempo limite de 3 segundos
	defer cancel()                                         // cancela o contexto após a função main terminar
	bookHotel(ctx)                                         // tenta reservar um hotel
}

func bookHotel(ctx context.Context) error {
	// simula um processo de reserva de hotel
	// que pode ser cancelado se um tempo limite for atingido
	select {
	case <-ctx.Done():
		fmt.Println("hotel booking cancelled, timeout reached")
		return fmt.Errorf("booking cancelled")
	case <-time.After(1 * time.Second):
		fmt.Println("hotel booking done")
		return nil
	}
}

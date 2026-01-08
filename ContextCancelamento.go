package main

import (
	"context"
	"time"
)

// ProcessNumbers processa uma lista de números respeitando cancelamento via context
func ProcessNumbers(ctx context.Context, numbers []int) ([]int, error) {
	var result []int

	for _, n := range numbers {
		select {
		case <-ctx.Done():
			// Retorno processado até o momento
			return result, ctx.Err()
		default:
			// Continua processamento
		}

		// Simula processamento
		time.Sleep(1 * time.Second)

		processed := n * 2
		result = append(result, processed)
	}

	return result, nil
}

package main

import (
	"errors"
	"net/http"
	"sync"
	"time"
)

// FetchURLs realiza requisições HTTP GET de forma concorrente
// limitando o número -> requisições simultâneas e aplicando timeout.

func FetchURLs(urls []string) (map[string]int, error) {
	if len(urls) == 0 {
		return nil, errors.New("lista de URLs vazia")
	}

	results := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Semáforo para limitar concorrência a 5
	sem := make(chan struct{}, 5)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var hasError bool

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			// Adquire vaga no semáforo
			sem <- struct{}{}
			defer func() { <-sem }()

			resp, err := client.Get(u)
			if err != nil {
				mu.Lock()
				results[u] = 0 // 0 indica falha na requisição
				hasError = true
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			mu.Lock()
			results[u] = resp.StatusCode
			mu.Unlock()
		}(url)
	}

	wg.Wait()

	if hasError {
		return results, errors.New("uma ou mais requisições falharam")
	}

	return results, nil

}

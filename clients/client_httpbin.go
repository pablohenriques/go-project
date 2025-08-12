package clients

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	config "github.com/pablohenriques/go-project/config"
	"github.com/sony/gobreaker"
)

func GetHttpBin() {
	cb := gobreaker.NewCircuitBreaker(config.GetSetting())
	client := resty.New()

	urlSuccess := fmt.Sprintf("https://httpbin.org/status%d", http.StatusOK)
	urlFailure := fmt.Sprintf("https://httpbin.org/status%d", http.StatusServiceUnavailable)

	for index := 0; index < 6; index++ {
		log.Printf("--- Tentativa %d ---", index+1)

		url := urlSuccess
		if index > 0 {
			url = urlFailure
		}

		resultado, err := cb.Execute(func() (interface{}, error) {
			return fazerRequisicao(client, url)
		})

		if err != nil {
			log.Printf("Resultado: Erro! -> %v\n", err)
		} else {
			log.Printf("Resultado: Sucesso! -> %v\n", resultado)
		}
		time.Sleep(time.Second)
	}

	log.Printf("\n--- Aguardando timeout Circuit Breaker (4 segundos) ---\n")
	time.Sleep(5 * time.Second)

	log.Println("--- Tentativa de recuperação (estado Meio-Aberto) ---")
	resultado, err := cb.Execute(func() (interface{}, error) {
		return fazerRequisicao(client, urlSuccess)
	})
	if err != nil {
		log.Printf("Resultado: Erro! -> %v\n", err)
	} else {
		log.Printf("Resultado: Sucesso! -> %v\n", resultado)
	}

}

func fazerRequisicao(client *resty.Client, url string) (string, error) {
	log.Printf(" -> Fazendo requisição para: %s", url)
	resp, err := client.R().Get(url)
	if err != nil {
		return "", err
	}
	if resp.IsError() {
		return "", fmt.Errorf("API retornou status de erro: %s", resp.Status())
	}
	return fmt.Sprintf("Sucesso com status: %s", resp.Status()), nil
}

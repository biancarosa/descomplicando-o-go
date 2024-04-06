package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Aula 2 - Arquivos e diretórios
// Aula 3 - Tipos básicos de variáveis, etc

// Aula 7 - Requisições web

type URL string
type HealthCheck struct {
	Url          URL
	StatusCode   int
	ResponseBody []byte
}

func (obj *HealthCheck) FormatResult() string {
	return fmt.Sprintf("%s - %d - %s", string(obj.Url), obj.StatusCode, string(obj.ResponseBody))
}

func main() {
	arquivo := "healthcheck"

	f, err := os.OpenFile(arquivo, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 1 defer vai ser ser executado no final da função atual
	// sempre coloco o mais proximo possivel da abertura do arquivo pra não esquecer
	// sempre depois de verificar se tem erro
	defer f.Close()

	url := "https://google.com.br/"
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	apiHealthCheck := &HealthCheck{
		ResponseBody: responseBytes,
		StatusCode:   resp.StatusCode,
		Url:          URL(url),
	}

	_, err = f.WriteString(apiHealthCheck.FormatResult())
	if err != nil {
		panic(err)
	}

}

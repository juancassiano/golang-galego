package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
	"time"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)

		duration := time.Since(start)
		log.Printf("Tempo de execução: %s", duration)
	}
}

// Autenticar verifica se usuário fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			respostas.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

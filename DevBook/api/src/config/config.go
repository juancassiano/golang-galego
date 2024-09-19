package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	SecretKey          []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 9000
	}
	StringConexaoBanco = "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	// StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_SENHA"),
	// 	os.Getenv("DB_NOME"),
	// )
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

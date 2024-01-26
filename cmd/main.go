package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/presentation"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	r := chi.NewRouter()

	presentation.PresentationList(r)

	fmt.Println("Aplication functionally.")
	http.ListenAndServe(":8083", r)
}

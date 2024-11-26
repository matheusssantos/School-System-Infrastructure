package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"subject-service/internal/infra/repository"
	"subject-service/internal/infra/web"
	"subject-service/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=host.docker.internal port=5432 dbname=scholl sslmode=disable")

	if err != nil {
		fmt.Println("Erro ao conectar-se ao Postgres")
		fmt.Print(err)
		return
	}

	defer db.Close() //Finaliza a conexão quando acabar

	registrationRepository := repository.NewRegistrationPostgres(db)

	createRegistrationUseCase := usecase.NewRegistrationUseCase(registrationRepository)

	registrationHandlers := web.NewRegistrationHandlers(createRegistrationUseCase)

	router := chi.NewRouter()
	router.Post("/registration", registrationHandlers.CreateRegistrationHandler)

	go http.ListenAndServe(":8080", router)
	print("Server on port 8080")
}

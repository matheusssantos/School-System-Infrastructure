package main

import (
	"database/sql"
	"log"
	"net/http"
	"subject-service/internal/infra/repository"
	"subject-service/internal/infra/web"
	"subject-service/internal/usecase"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql" // Importando o driver MySQL
	_ "github.com/lib/pq"              // Importando o driver PostgreSQL
)

func main() {
	// Conexão com o banco de dados do school (PostgreSQL)
	log.Print("Tentando conectador com o banco")

	dbSchool, err := sql.Open(
		"postgres",
		"user=postgres password=postgres host=postgres port=5432 dbname=school sslmode=disable",
	)

	if err != nil {
		log.Println("Erro ao conectar-se ao banco school")
		log.Print(err)
		return
	}
	log.Println("Conectado ao banco users")

	defer dbSchool.Close() // Fechar a conexão quando terminar

	// Conexão com o banco de dados de users (MySQL)
	//dbUsers, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/users")

	dbUsers, err := sql.Open(
		"mysql",
		"root:root@tcp(mysql:3306)/user",
	)

	if err != nil {
		log.Println("Erro ao conectar-se ao banco users")
		log.Print(err)
		return
	}

	log.Println("Conectado ao banco users")
	defer dbUsers.Close()

	// Criação do repositório de registros, passando as conexões para os dois bancos de dados
	registrationRepository := repository.NewRegistrationPostgres(dbSchool, dbUsers)

	// Criar a tabela 'registration' no banco school (PostgreSQL)
	// err = registrationRepository.CreateRegistrationTable(dbSchool)
	// if err != nil {
	// 	log.Fatal("Erro ao criar a tabela:", err)
	// }

	// Criação do caso de uso para a matrícula
	createRegistrationUseCase := usecase.NewRegistrationUseCase(registrationRepository)

	// Criação do caso de uso para consultas
	registrationQueriesUseCase := usecase.NewRegistrationQueriesUseCase(registrationRepository)

	// Criando os handlers
	registrationHandlers := web.NewRegistrationHandlers(createRegistrationUseCase, registrationQueriesUseCase)

	// Configuração do roteador HTTP
	router := chi.NewRouter()

	// Definindo as rotas
	router.Post("/registration", registrationHandlers.CreateRegistrationHandler)
	router.Post("/find-groups-by-student", registrationHandlers.FindGroupsByStudentIDHandler)
	router.Post("/find-students-by-group", registrationHandlers.FindStudentsByGroupIDHandler)

	/// Mensagem de sucesso
	log.Println("Server on port 8080")
	http.ListenAndServe(":8080", router)

}

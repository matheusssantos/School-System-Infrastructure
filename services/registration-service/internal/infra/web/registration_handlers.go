package web

import (
	"encoding/json"
	"net/http"
	"subject-service/internal/usecase"
)

type RegistrationHandlers struct {
	CreateRegistrationUseCase  *usecase.CreateRegistrationUseCase
	RegistrationQueriesUseCase *usecase.RegistrationQueriesUseCase
	//Others handlers
}

func NewRegistrationHandlers(createRegistrationUseCase *usecase.CreateRegistrationUseCase, registrationQueriesUseCase *usecase.RegistrationQueriesUseCase) *RegistrationHandlers {
	return &RegistrationHandlers{
		CreateRegistrationUseCase:  createRegistrationUseCase,
		RegistrationQueriesUseCase: registrationQueriesUseCase,
	}
}

func (r *RegistrationHandlers) CreateRegistrationHandler(w http.ResponseWriter, req *http.Request) {
	var input usecase.CreateRegistrationInputDto
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := r.CreateRegistrationUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// Handler para consultar as turmas em que um estudante está matriculado
func (r *RegistrationHandlers) FindGroupsByStudentIDHandler(w http.ResponseWriter, req *http.Request) {
	var input usecase.FindGroupsByStudentIDInputDto
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := r.RegistrationQueriesUseCase.FindGroupsByStudentID(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// Handler para consultar os estudantes matriculados em uma turma
func (r *RegistrationHandlers) FindStudentsByGroupIDHandler(w http.ResponseWriter, req *http.Request) {
	var input usecase.FindStudentsByGroupIDInputDto
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := r.RegistrationQueriesUseCase.FindStudentsByGroupID(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

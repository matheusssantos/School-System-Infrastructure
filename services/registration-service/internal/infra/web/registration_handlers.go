package web

import (
	"encoding/json"
	"net/http"
	"subject-service/internal/usecase"
)

type RegistrationHandlers struct {
	CreateRegistrationUseCase *usecase.CreateRegistrationUseCase
	//Others handlers
}

func NewRegistrationHandlers(createRegistrationUseCase *usecase.CreateRegistrationUseCase) *RegistrationHandlers {
	return &RegistrationHandlers{
		CreateRegistrationUseCase: createRegistrationUseCase,
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

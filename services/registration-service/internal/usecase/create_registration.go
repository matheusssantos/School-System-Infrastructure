package usecase

import (
	"log"
	"subject-service/internal/entity"
)

type CreateRegistrationInputDto struct {
	UserID  int `json:"user_id"`
	GroupID int `json:"group_id"`
}

type CreateRegistrationOutputDto struct {
	UserID  int
	GroupID int
}

type CreateRegistrationUseCase struct {
	ResgistrationRepository entity.ResgistrationRepository
}

func NewRegistrationUseCase(registrationRepository entity.ResgistrationRepository) *CreateRegistrationUseCase {
	return &CreateRegistrationUseCase{
		ResgistrationRepository: registrationRepository,
	}
}

func (uc *CreateRegistrationUseCase) Execute(input CreateRegistrationInputDto) (*CreateRegistrationOutputDto, error) {
	registration := entity.NewRegistration(input.UserID, input.GroupID)
	err := uc.ResgistrationRepository.Create(registration)
	if err != nil {
		log.Println("Erro ao criar matrícula")
		log.Println(err)
		return nil, err
	}

	return &CreateRegistrationOutputDto{
		UserID:  registration.UserID,
		GroupID: registration.GroupID,
	}, nil
}

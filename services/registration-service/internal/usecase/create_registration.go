package usecase

import "subject-service/internal/entity"

type CreateRegistrationInputDto struct {
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`
}

type CreateRegistrationOutputDto struct {
	UserID  string
	GroupID string
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
		return nil, err
	}

	return &CreateRegistrationOutputDto{
		UserID:  registration.UserID,
		GroupID: registration.GroupID,
	}, nil
}

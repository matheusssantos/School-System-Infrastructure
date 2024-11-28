package usecase

import (
	"subject-service/internal/entity"
)

type RegistrationQueriesUseCase struct {
	ResgistrationRepository entity.ResgistrationRepository
}

func NewRegistrationQueriesUseCase(registrationRepository entity.ResgistrationRepository) *RegistrationQueriesUseCase {
	return &RegistrationQueriesUseCase{
		ResgistrationRepository: registrationRepository,
	}
}

// Input e Output para "Consultar Turmas de um Estudante"
type FindGroupsByStudentIDInputDto struct {
	StudentID int `json:"user_id"`
}

type GroupDto struct {
	ID           int    `json:"id"`
	DisciplineID int    `json:"discipline_id"`
	Name         string `json:"name"`
}

type FindGroupsByStudentIDOutputDto struct {
	Groups []GroupDto `json:"groups"`
}

func (uc *RegistrationQueriesUseCase) FindGroupsByStudentID(input FindGroupsByStudentIDInputDto) (*FindGroupsByStudentIDOutputDto, error) {
	groups, err := uc.ResgistrationRepository.FindGroupsByStudentID(input.StudentID)
	if err != nil {
		return nil, err
	}

	// Transformar os resultados em DTOs
	groupDtos := make([]GroupDto, len(groups))
	for i, group := range groups {
		groupDtos[i] = GroupDto{
			ID:           group.ID,
			DisciplineID: group.DisciplineID,
		}
	}

	return &FindGroupsByStudentIDOutputDto{
		Groups: groupDtos,
	}, nil
}

// Input e Output para "Consultar Estudantes de uma Turma"
type FindStudentsByGroupIDInputDto struct {
	GroupID int `json:"group_id"`
}

type StudentDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	RG   string `json:"rg"`
	Type string `json:"type"`
}

type FindStudentsByGroupIDOutputDto struct {
	Students []StudentDto `json:"students"`
}

func (uc *RegistrationQueriesUseCase) FindStudentsByGroupID(input FindStudentsByGroupIDInputDto) (*FindStudentsByGroupIDOutputDto, error) {
	students, err := uc.ResgistrationRepository.FindStudentsByGroupID(input.GroupID)
	if err != nil {
		return nil, err
	}

	// Transformar os resultados em DTOs
	studentDtos := make([]StudentDto, len(students))
	for i, student := range students {
		studentDtos[i] = StudentDto{
			ID:   student.ID,
			Name: student.Name,
			RG:   student.RG,
			Type: student.Type,
		}
	}

	return &FindStudentsByGroupIDOutputDto{
		Students: studentDtos,
	}, nil
}

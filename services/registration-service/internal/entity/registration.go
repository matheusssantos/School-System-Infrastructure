package entity

type ResgistrationRepository interface {
	// Criação de matrícula
	Create(registration *Registration) error

	// Consultar turmas de um estudante
	FindGroupsByStudentID(studentID int) ([]Group, error)

	// Consultar estudantes de uma turma
	FindStudentsByGroupID(groupID int) ([]Student, error)
}

type Registration struct {
	UserID  string
	GroupID string
}

func NewRegistration(UserID string, GroupID string) *Registration {
	return &Registration{
		UserID:  UserID,
		GroupID: GroupID,
	}
}

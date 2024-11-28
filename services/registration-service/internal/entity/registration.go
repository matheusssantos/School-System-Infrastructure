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
	UserID  int
	GroupID int
}

func NewRegistration(UserID int, GroupID int) *Registration {
	return &Registration{
		UserID:  UserID,
		GroupID: GroupID,
	}
}

package entity

type Student struct {
	ID      int     // ID único do estudante
	Name    string  // Nome completo do estudante
	RG      string  // RG do estudante
	Type    string  // Tipo de usuário (ex.: "Aluno")
	Address Address // Endereço do estudante
}

type StudentRepository interface {
	Create(student *Student) error
	FindByID(id int) (*Student, error)
	FindAll() ([]Student, error)
}

func NewStudent(id int, name, rg, userType string, address Address) *Student {
	return &Student{
		ID:   id,
		Name: name,
		RG:   rg,
		Type: userType,
	}
}

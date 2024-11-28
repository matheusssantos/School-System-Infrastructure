package entity

type Group struct {
	ID           int // ID único do grupo
	DisciplineID int // ID da disciplina associada ao grupo
}

type GroupRepository interface {
	Create(group *Group) error
	FindByID(id int) (*Group, error)
}

func NewGroup(id int, disciplineID int) *Group {
	return &Group{
		ID:           id,
		DisciplineID: disciplineID,
	}
}

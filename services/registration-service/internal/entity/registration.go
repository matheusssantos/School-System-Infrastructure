package entity

type ResgistrationRepository interface {
	Create(registration *Registration) error
	// FindAll() ([]Registration, error)
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

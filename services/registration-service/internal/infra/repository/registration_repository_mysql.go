package repository

import (
	"database/sql"
	"subject-service/internal/entity"
)

type RegistrationRepository struct {
	DB *sql.DB
}

// FindAll implements entity.ResgistrationRepository.
func (r *RegistrationRepository) FindAll() ([]entity.Registration, error) {
	panic("unimplemented")
}

func NewRegistrationPostgres(db *sql.DB) *RegistrationRepository {
	return &RegistrationRepository{
		DB: db,
	}
}

func (r *RegistrationRepository) Create(registration *entity.Registration) error {
	_, err := r.DB.Exec("INSERT INTO registration (user_id, group_id) VALUES (?, ?)", registration.UserID, registration.GroupID)

	if err != nil {
		return err
	}

	return nil
}

// func (r *RegistrationRepository) FindAll() ([]*entity.Registration, error) {
// 	rows, err := r.DB.Query("SELECT user_id, group_id FROM registration")

// 	if err != nil {
// 		return nil, err
// 	}

// 	var registrations []*entity.Registration
// 	for rows.Next() {
// 		var registration entity.Registration
// 		err = rows.Scan(&registration.UserID, &registration.GroupID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		registrations = append(registrations, &registration)
// 	}

// 	return registrations, nil
// }

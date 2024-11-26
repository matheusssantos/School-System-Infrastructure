package repository

import (
	"database/sql"
	"fmt"
	"subject-service/internal/entity"
)

type RegistrationRepository struct {
	DBSchool *sql.DB
	DBUsers  *sql.DB
}

// FindAll implements entity.ResgistrationRepository.
func (r *RegistrationRepository) FindAll() ([]entity.Registration, error) {
	panic("unimplemented")
}

func NewRegistrationPostgres(dbSchool *sql.DB, dbUsers *sql.DB) *RegistrationRepository {
	return &RegistrationRepository{
		DBSchool: dbSchool,
		DBUsers:  dbUsers,
	}
}

// Criação da matricula
func (r *RegistrationRepository) Create(registration *entity.Registration) error {
	// Verificar se o estudante existe no banco de usuários
	var studentName string
	err := r.DBUsers.QueryRow("SELECT name FROM users WHERE id = ?", registration.UserID).Scan(&studentName)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("estudante com ID %d não encontrado", registration.UserID)
		}
		return err
	}

	// Verificar se a turma existe no banco de disciplinas/turmas
	var groupName string
	err = r.DBSchool.QueryRow("SELECT name FROM groups WHERE id = ?", registration.GroupID).Scan(&groupName)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("turma com ID %d não encontrada", registration.GroupID)
		}
		return err
	}

	// Matricular o estudante na turma
	_, err = r.DBSchool.Exec(
		"INSERT INTO registration (user_id, group_id) VALUES (?, ?)",
		registration.UserID, registration.GroupID,
	)
	if err != nil {
		return fmt.Errorf("erro ao matricular estudante na turma: %v", err)
	}

	return nil
}

func (r *RegistrationRepository) FindGroupsByStudentID(studentID int) ([]entity.Group, error) {
	query := `
		SELECT g.id, g.discipline_id, g.name
		FROM registration r
		JOIN groups g ON r.group_id = g.id
		WHERE r.user_id = ?
	`
	rows, err := r.DBSchool.Query(query, studentID)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar turmas para o estudante: %v", err)
	}
	defer rows.Close()

	var groups []entity.Group
	for rows.Next() {
		var group entity.Group
		if err := rows.Scan(&group.ID, &group.DisciplineID, &group.Name); err != nil {
			return nil, fmt.Errorf("erro ao escanear resultados: %v", err)
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro no processamento das linhas: %v", err)
	}

	return groups, nil
}

func (r *RegistrationRepository) FindStudentsByGroupID(groupID int) ([]entity.Student, error) {
	// Consultar os IDs dos estudantes matriculados na turma
	query := `
		SELECT user_id
		FROM registration
		WHERE group_id = ?
	`
	rows, err := r.DBSchool.Query(query, groupID)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar estudantes matriculados: %v", err)
	}
	defer rows.Close()

	var studentIDs []int
	for rows.Next() {
		var userID int
		if err := rows.Scan(&userID); err != nil {
			return nil, fmt.Errorf("erro ao escanear user_id: %v", err)
		}
		studentIDs = append(studentIDs, userID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro no processamento das linhas: %v", err)
	}

	// Buscar informações dos estudantes no DBUsers
	var students []entity.Student
	for _, id := range studentIDs {
		var student entity.Student
		err := r.DBUsers.QueryRow("SELECT id, name, rg, type FROM users WHERE id = ?", id).
			Scan(&student.ID, &student.Name, &student.RG, &student.Type)
		if err != nil {
			return nil, fmt.Errorf("erro ao buscar dados do estudante ID %d: %v", id, err)
		}
		students = append(students, student)
	}

	return students, nil
}

func (r *RegistrationRepository) CreateRegistrationTable(db *sql.DB) error {
	// Comando SQL para criar a tabela se não existir, sem chave estrangeira para user_id
	query := `
	CREATE TABLE IF NOT EXISTS registration (
		user_id VARCHAR(255) NOT NULL,
		group_id VARCHAR(255) NOT NULL,
		PRIMARY KEY (user_id, group_id),
		CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
	);
	`

	// Executando o comando no banco de dados
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao criar a tabela registration: %v", err)
	}

	fmt.Println("Tabela 'registration' criada ou já existe.")
	return nil
}

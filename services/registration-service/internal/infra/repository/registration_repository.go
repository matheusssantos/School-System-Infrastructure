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

// Criação da matrícula
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

	// Verificar se o grupo (turma) existe no banco de grupos
	var groupExists int
	err = r.DBSchool.QueryRow("SELECT 1 FROM groups WHERE id = ?", registration.GroupID).Scan(&groupExists)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("grupo com ID %d não encontrado", registration.GroupID)
		}
		return err
	}

	// Verificar se o estudante já está matriculado no grupo
	var existingRegistration int
	err = r.DBSchool.QueryRow("SELECT 1 FROM registration WHERE user_id = ? AND group_id = ?", registration.UserID, registration.GroupID).Scan(&existingRegistration)
	if err != nil {
		if err == sql.ErrNoRows {
			// Nenhuma matrícula encontrada, podemos prosseguir para criar a matrícula
		} else {
			return err
		}
	} else {
		// Se o valor retornado for 1, significa que o estudante já está matriculado
		return fmt.Errorf("estudante com ID %d já está matriculado no grupo com ID %d", registration.UserID, registration.GroupID)
	}

	// Criar a matrícula (registrar o estudante na turma)
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
		SELECT g.id, g.subject_id
		FROM registrations r
		JOIN group g ON r.id = g.subject_id
		WHERE g.user_id = ?
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
		FROM group
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
		err := r.DBUsers.QueryRow("SELECT id, name, rg, type FROM user WHERE id = ?", id).
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

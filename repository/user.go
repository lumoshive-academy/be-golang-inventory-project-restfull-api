package repository

import (
	"database/sql"
	"go-25-27/model"

	"go.uber.org/zap"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindAllStudents() ([]model.User, error)
	GetUserByID(id int) (model.User, error)
}

type userRepositoryImpl struct {
	DB  *sql.DB
	Log *zap.Logger
}

func NewUserRepository(db *sql.DB, log *zap.Logger) UserRepository {
	return &userRepositoryImpl{
		DB:  db,
		Log: log,
	}
}

func (r *userRepositoryImpl) Create(user *model.User) error {
	query := `
		INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id
	`
	return r.DB.QueryRow(query, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID)
}

func (r *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, created_at, updated_at, deleted_at, name, email, password, role
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`
	var user model.User
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
		&user.Name, &user.Email, &user.Password, &user.Role,
	)
	// if err == sql.ErrNoRows {
	// 	return nil, nil // user tidak ditemukan
	// }
	r.Log.Info("Repository User", zap.String("query", query))
	return &user, err
}

func (r *userRepositoryImpl) FindAllStudents() ([]model.User, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, password, role FROM users WHERE role = 'student'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role)
		if err != nil {
			return nil, err
		}
		students = append(students, u)
	}
	return students, nil
}

func (r *userRepositoryImpl) GetUserByID(id int) (model.User, error) {
	var user model.User
	query := "SELECT id, name, email, role FROM users WHERE id = $1"

	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Role)
	if err != nil {
		return user, err
	}

	return user, nil
}

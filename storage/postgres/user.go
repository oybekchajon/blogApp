package postgres

import (
	"database/sql"
	"time"

)

type User struct {
	ID              int64
	FirstName       string
	LastName        string
	PhoneNumber     string
	Email           string
	Gender          string
	Password        string
	Username        string
	ProfileImageUrl string
	Type            string
	CreatedAt       time.Time
}

type GetAllUsersParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetBookResult struct {
	Users []*User
	Count int32
}

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db}
}

func (m *DBManager) CreateUser(user *User) (*User, error){
	query := `
		INSERT INTO users(
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			username,
			profile_image_url,
			type
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at
	`

	row := m.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.Gender,
		user.Password,
		user.Username,
		user.ProfileImageUrl,
		user.Type,
	)

	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *DBManager) GetUser(id int) (*User, error){
	query := `
		SELECT * FROM users WHERE id=$1
	`

	row := m.db.QueryRow(query,id)

	var result User

	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.Gender,
		&result.Password,
		&result.Username,
		&result.ProfileImageUrl,
		&result.Type,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


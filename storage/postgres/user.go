package postgres

import (
	"database/sql"
	"fmt"
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

func (m *DBManager) CreateUser(user *User) (*User, error) {
	query := `
		INSERT INTO users(
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			user_name,
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

func (m *DBManager) GetUser(id int) (*User, error) {
	query := `
		SELECT 
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			user_name,
			profile_image_url,
			type
		FROM users WHERE id=$1
	`

	row := m.db.QueryRow(query, id)

	var result User

	err := row.Scan(
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

func (m *DBManager) UpdateUser(user *User) (*User, error) {
	query := `
		UPDATE users SET
			first_name=$1,
			last_name=$2,
			phone_number=$3,
			email=$4,
			gender=$5,
			password=$6,
			user_name=$7,
			profile_image_url=$8,
			type=$9
		WHERE id=$10
		RETURNING first_name, last_name, phone_number, email, gender, password, user_name, profile_image_url, type
	`

	row := m.db.QueryRow(query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.Gender,
		user.Password,
		user.Username,
		user.ProfileImageUrl,
		user.Type,
		user.ID,
	)

	var result User

	err := row.Scan(
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

	return &result, err
}

func (m *DBManager) GetAll(params *GetAllUsersParams) (*GetBookResult, error) {
	result := GetBookResult{
		Users: make([]*User, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s' 
				OR username ILIKE '%s' OR phone_number ILIKE '%s'`,
			str, str, str, str, str,
		)
	}

	query := `
		SELECT 
			id,
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			user_name,
			profile_image_url,
			type,
			created_at
		FROM users
		` + filter +`
		ORDER BY created_at desc
		` + limit

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var u User

		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.PhoneNumber,
			&u.Email,
			&u.Gender,
			&u.Password,
			&u.Username,
			&u.ProfileImageUrl,
			&u.Type,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result.Users = append(result.Users, &u)
	}
	return &result, nil
}

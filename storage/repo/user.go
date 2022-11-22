package repo

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int64
	FirstName       string
	LastName        string
	PhoneNumber     *string
	Email           string
	Gender          *string
	Password        string
	Username        string
	ProfileImageUrl *string
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

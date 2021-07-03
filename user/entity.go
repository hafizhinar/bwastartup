package user

import "time"

type User struct {
	ID           uint64
	Name         string
	Occupation   string
	Email        string
	PasswordHash string
	Profile      string
	Role         string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

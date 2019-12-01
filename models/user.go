package models

import (
	"github.com/gobuffalo/uuid"
	"time"
)

var userTypes = map[string]int{"student": 1, "staff": 2, "parent": 3}

type User struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Type      int       `db:"type"`
	Email     string    `db:"email"`
	Password  string    `db:"password" rw:"w"`
}

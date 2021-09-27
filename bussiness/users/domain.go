package users

import "time"

type User struct {
	ID        int
	NIK       int
	Password  string
	FotoRumah string
	FotoDiri  string
	Alamat    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Functions interface {
	Created(user *User) (*User, error)
	Update(user *User, id int) (*User, error)
	FindByID(id int) (*User, error)
	Deleted(id int) (*User, error)
}

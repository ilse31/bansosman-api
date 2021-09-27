package users

import "time"

type Users struct {
	ID        int
	NIK       int
	Password  string
	FotoRumah string
	FotoDiri  string
	Alamat    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

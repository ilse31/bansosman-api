package request

import (
	"bansosman/bussiness/admin"
)

type Admins struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Admins) ToDomain() *admin.Domain {
	return &admin.Domain{
		Username: req.Username,
		Password: req.Password,
		RoleID:   req.RoleID,
	}
}

package entity

type RequestAuth struct {
	Nik      string `json:"nik" validate:"required"`
	Password string `json:"password" validate:"required"`
}

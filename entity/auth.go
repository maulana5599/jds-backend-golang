package entity

type RequestAuth struct {
	Nik      string `json:"nik" validate:"required,min=1,max=16"`
	Password string `json:"password" validate:"required"`
}

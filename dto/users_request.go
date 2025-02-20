package dto

type UserRegistration struct {
	Name          string `json:"nama"  validate:"required"`
	NationalityID string `json:"nik" validate:"required"`
	PhoneNumber   string `json:"no_hp" validate:"required"`
}

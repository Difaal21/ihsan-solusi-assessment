package dto

type UserTOTPVerification struct {
	Code string `json:"code" validate:"required,min=6"`
}

type UserRegistration struct {
	Name          string `json:"nama"  validate:"required"`
	NationalityID string `json:"nik" validate:"required"`
	PhoneNumber   string `json:"no_hp" validate:"required"`
}

type Credit struct {
	BankAccountNumber string  `json:"no_rekening" validate:"required"`
	Amount            float64 `json:"nominal" validate:"required"`
}

type Debit struct {
	BankAccountNumber string  `json:"no_rekening" validate:"required"`
	Amount            float64 `json:"nominal" validate:"required"`
}

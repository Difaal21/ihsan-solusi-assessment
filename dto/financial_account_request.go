package dto

type Credit struct {
	BankAccountNumber string  `json:"no_rekening" validate:"required"`
	Amount            float64 `json:"nominal" validate:"required"`
}

type Debit struct {
	BankAccountNumber string  `json:"no_rekening" validate:"required"`
	Amount            float64 `json:"nominal" validate:"required"`
}

type CheckBalance struct {
	BankAccountNumber string `json:"no_rekening" validate:"required"`
}

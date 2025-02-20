package dto

type UserRegistrationReponse struct {
	BankAccountNumber string `json:"no_rekening"`
}

func NewUserRegistrationReponse(bankAccountNumber string) UserRegistrationReponse {
	return UserRegistrationReponse{
		BankAccountNumber: bankAccountNumber,
	}
}

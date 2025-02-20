package dto

type CheckBalanceResponse struct {
	AccountNumber string  `json:"no_rekening"`
	Balance       float64 `json:"saldo"`
}

func NewCheckBalanceResponse(accountNumber string, balance float64) (result *CheckBalanceResponse) {
	return &CheckBalanceResponse{
		AccountNumber: accountNumber,
		Balance:       balance,
	}
}

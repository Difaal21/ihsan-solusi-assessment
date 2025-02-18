package responses

type Responses interface {
	SetData(data interface{}) Responses
	SetMessage(message string) Responses
	SetPagination(pagination interface{}) Responses
	Send() *ResponsesImpl
}

type ResponsesImpl struct {
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
	Status     string      `json:"status"`
}

func NewResponse() *ResponsesImpl {
	return &ResponsesImpl{}
}

package resource

type GeneralResource struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

func NewGeneralResource(code int16, message string) *GeneralResource {
	return &GeneralResource{Code: code, Message: message}
}

package product

type ProductResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponse() {
}

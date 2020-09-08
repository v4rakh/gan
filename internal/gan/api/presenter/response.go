package presenter

const (
	ErrorBadRequest = "BadRequest"
	ErrorNotFound   = "NotFound"
	ErrorGeneral    = "GeneralError"
)

type Response struct {
}

type DataResponse struct {
	Response
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status string `json:"status,omitempty"`
	DataResponse
}

func NewDataResponseWithPayload(payload interface{}) *DataResponse {
	e := new(DataResponse)
	e.Data = payload
	return e
}

func NewErrorResponseWithMessage(message string) *ErrorResponse {
	e := new(ErrorResponse)
	e.Status = ErrorGeneral
	e.Message = message
	return e
}

func NewErrorResponseWithStatusAndMessage(status string, message string) *ErrorResponse {
	e := new(ErrorResponse)
	e.Status = status
	e.Message = message
	return e
}

func NewErrorResponseWithStatusAndMessageAndData(status string, message string, data interface{}) *ErrorResponse {
	e := new(ErrorResponse)
	e.Status = status
	e.Message = message
	e.Data = data
	return e
}

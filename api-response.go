package go_commons

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
}

type Responder interface {
	SetData(data interface{}) *ApiResponse
	SetMessage(message string) *ApiResponse
	SetStatus(status bool) *ApiResponse
	SetCode(code int) *ApiResponse
}

const (
	successMessage = "succeed"
	failedMessage  = "failed"
	successStatus  = true
	failedStatus   = false
	successCode    = 200
	failedCode     = 500
)

func NewApiResponse(message string, status bool, code int, data interface{}) *ApiResponse {
	return &ApiResponse{Message: message, Data: data, Status: status, Code: code}
}

func NewSuccessfulResponse(data interface{}) *ApiResponse {
	return NewApiResponse(successMessage, successStatus, successCode, data)
}

func NewFailedResponse(message string) *ApiResponse {
	return NewApiResponse(message, failedStatus, failedCode, nil)
}

func (response *ApiResponse) SetData(data interface{}) *ApiResponse {
	response.Data = data
	return response
}

func (response *ApiResponse) SetMessage(message string) *ApiResponse {
	response.Message = message
	return response
}

func (response *ApiResponse) SetStatus(status bool) *ApiResponse {
	response.Status = status
	return response
}

func (response *ApiResponse) SetCode(code int) *ApiResponse {
	response.Code = code
	return response
}

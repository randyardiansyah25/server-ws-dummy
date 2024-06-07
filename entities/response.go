package entities

type BaseResponse struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

type BaseResponseString struct {
	BaseResponse
	ResponseData string `json:"response_data"`
}

type BaseResponseAny struct {
	BaseResponse
	ResponseData any `json:"response_data"`
}

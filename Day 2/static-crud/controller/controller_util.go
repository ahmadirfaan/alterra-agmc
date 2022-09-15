package controller

import "static-crud/model"

func wrapperResponse(code int, message string, response interface{}) *model.HTTPResponse {
	newResponse := new(model.HTTPResponse)
	newResponse.Code = code
	newResponse.Message = message
	newResponse.Data = response
	return newResponse
}

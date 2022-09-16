package controller

import (
	models "alterra-agmc-day3/models/website"
)

func wrapperResponse(code int, message string, response interface{}) *models.HTTPResponse {
	newResponse := new(models.HTTPResponse)
	newResponse.Code = code
	newResponse.Message = message
	newResponse.Data = response
	return newResponse
}

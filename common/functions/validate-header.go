package functions

import (
	"glow-service/models"
	"net/http"
)

const (
	platformHeaderParamName      = "platform"
	authorizationHeaderParamName = "authorization"
)

/*
	Extracts request header params and return an object.
	The same object has a stackTrace log to use
*/
func ValidateHeader(header *http.Header) *models.StackLog {
	var log models.StackLog
	log.Platform = header.Get(platformHeaderParamName)
	log.User.Token = header.Get(authorizationHeaderParamName)

	return &log
}

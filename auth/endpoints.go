package auth

import (
	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/auth/login/",
		Handler: loginHandler,
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/auth/reset_password/",
		Handler: resetPasswordHandler,
	},
}

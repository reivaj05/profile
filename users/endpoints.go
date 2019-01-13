package users

import (
	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/users/{id:[0-9]+}/",
		Handler: getItemHandler,
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/users/",
		Handler: postItemhandler,
	},
	&GoServer.Endpoint{
		Method:  "PUT",
		Path:    "/users/{id:[0-9]+}/",
		Handler: putItemHandler,
	},
}

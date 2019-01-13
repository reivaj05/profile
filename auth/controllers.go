package auth

import (
	"fmt"
	"net/http"

	"github.com/reivaj05/GoJSON"
	"github.com/reivaj05/GoServer"
)

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	data, err := getJSONData(req)
	if err != nil {
		sendBadRequestResponse(rw)
		return
	}
	if !hasDataCredentials(data) {
		sendBadRequestResponse(rw)
		return
	}
	if user, err := login(data); err != nil {
		GoServer.SendResponseWithStatus(rw, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
	} else {
		GoServer.SendResponseWithStatus(rw, user.ToJSON(), http.StatusOK)
	}
}

func hasDataCredentials(data *GoJSON.JSONWrapper) bool {
	return data.HasPath("email") && data.HasPath("password")
}

func resetPasswordHandler(rw http.ResponseWriter, req *http.Request) {
	data, err := getJSONData(req)
	if err != nil {
		sendBadRequestResponse(rw)
		return
	}
	if !data.HasPath("email") {
		sendBadRequestResponse(rw)
		return
	}
	if err := resetPassword(data); err != nil {
		GoServer.SendResponseWithStatus(rw, fmt.Sprintf(`{"error": "%s" }`, err.Error()), http.StatusBadRequest)
	} else {
		GoServer.SendResponseWithStatus(rw, "", http.StatusOK)
	}
}

func getJSONData(req *http.Request) (data *GoJSON.JSONWrapper, err error) {
	body, _ := GoServer.ReadBodyRequest(req)
	data, err = GoJSON.New(body)
	return
}

func sendBadRequestResponse(rw http.ResponseWriter) {
	GoServer.SendResponseWithStatus(rw, GoServer.BadRequest, http.StatusBadRequest)
}

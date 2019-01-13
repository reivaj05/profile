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
	if err := login(data); err != nil {
		GoServer.SendResponseWithStatus(rw, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
	} else {
		GoServer.SendResponseWithStatus(rw, "", http.StatusOK)
	}
}

func hasDataCredentials(data *GoJSON.JSONWrapper) bool {
	return data.HasPath("email") && data.HasPath("password")
}

func logoutHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement logout")
	if err := logout(); err != nil {
		// GoServer.SendResponseWithStatus(rw, "", http.StatusInternalServerError)
	} else {
		GoServer.SendResponseWithStatus(rw, "success", http.StatusOK)
	}
}

func resetPasswordHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement reset password")
	if err := resetPassword(); err != nil {
		// GoServer.SendResponseWithStatus(rw, "", http.StatusInternalServerError)
	} else {
		GoServer.SendResponseWithStatus(rw, "success", http.StatusOK)
	}
}

func signupHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement signup")
	data, err := getJSONData(req)
	if err != nil {
		// GoServer.SendResponseWithStatus(rw, "Error reading data", http.StatusBadRequest)
		return
	}
	if err := validateSignupData(data); err != nil {
		// GoServer.SendResponseWithStatus(rw, "", http.StatusBadRequest)
		return
	}
	if err := signup(data); err != nil {
		// GoServer.SendResponseWithStatus(rw, "", http.StatusInternalServerError)
	} else {
		GoServer.SendResponseWithStatus(rw, "success", http.StatusOK)
	}
}

func validateSignupData(data *GoJSON.JSONWrapper) error {
	return nil
}

func getJSONData(req *http.Request) (data *GoJSON.JSONWrapper, err error) {
	body, _ := GoServer.ReadBodyRequest(req)
	data, err = GoJSON.New(body)
	return
}

func isDataValid(data *GoJSON.JSONWrapper) bool {
	// TODO: Implement your own logic
	return true
}

func sendBadRequestResponse(rw http.ResponseWriter) {
	GoServer.SendResponseWithStatus(rw, GoServer.BadRequest, http.StatusBadRequest)
}

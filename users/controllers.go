package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/reivaj05/GoJSON"
	"github.com/reivaj05/GoServer"
)

func getItemHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement get")
	params := GoServer.GetQueryParams(req)
	id, _ := strconv.Atoi(params["id"])
	if user, err := getUser(id); err != nil {
		GoServer.SendResponseWithStatus(rw, "error", http.StatusNotFound)
	} else {
		GoServer.SendResponseWithStatus(rw, user.toJSON(), http.StatusOK)
	}
}

func postItemhandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement post")
	data, err := getJSONData(req)
	if err != nil {
		sendBadRequestResponse(rw)
		return
	}
	if !isDataValid(data) {
		sendBadRequestResponse(rw)
		return
	}
	if user, err := newUser(data); err != nil {
		GoServer.SendResponseWithStatus(rw, "error", http.StatusInternalServerError)
	} else {
		GoServer.SendResponseWithStatus(rw, user.toJSON(), http.StatusCreated)
	}
}

func putItemHandler(rw http.ResponseWriter, req *http.Request) {
	data, err := getJSONData(req)
	if err != nil {
		sendBadRequestResponse(rw)
		return
	}
	params := GoServer.GetQueryParams(req)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("TODO: Implement put", id)
	if !isDataValid(data) {
		sendBadRequestResponse(rw)
		return
	}
	if user, err := getUser(id); err != nil {
		GoServer.SendResponseWithStatus(rw, "error", http.StatusNotFound)
	} else {
		updateUser(rw, user, data)
	}
}

func updateUser(rw http.ResponseWriter, user *User, data *GoJSON.JSONWrapper) {
	if err := user.update(data); err != nil {
		GoServer.SendResponseWithStatus(rw, "error", http.StatusInternalServerError)
	} else {
		GoServer.SendResponseWithStatus(rw, user.toJSON(), http.StatusOK)
	}
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
	GoServer.SendResponseWithStatus(
		rw, GoServer.BadRequest, http.StatusBadRequest)
}

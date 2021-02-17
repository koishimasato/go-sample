package handler

import (
	"encoding/json"
	"github.com/koishimasato/go-sample/internal/application/command"
	"github.com/koishimasato/go-sample/internal/application/service"
	"github.com/koishimasato/go-sample/internal/domain/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type UserHandler interface {
	Handle(http.ResponseWriter, *http.Request)
	Index(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type UserHandlerImpl struct {
	userApplicationService service.UserApplicationService
}

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"userType"`
}

type IndexResponse struct {
	Users []UserResponse `json:"users"`
}

type PostRequest struct {
	Name string `json:"name"`
}

type PutRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserHandler(s service.UserApplicationService) UserHandler {
	return &UserHandlerImpl{userApplicationService: s}
}

func (h *UserHandlerImpl) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		if strings.HasPrefix(r.URL.Path, "/users/") {
			h.Get(w, r)
		} else {
			h.Index(w, r)
		}
	case "POST":
		h.Post(w, r)
	case "PUT", "PATCH":
		h.Put(w, r)
	case "DELETE":
		h.Delete(w, r)
	}
}

func (h *UserHandlerImpl) Index(w http.ResponseWriter, _ *http.Request) {
	result, err := h.userApplicationService.GetAll()
	if err != nil {
		h.handleError(w, err)
		return
	}

	res := new(IndexResponse)
	for _, user := range result.Users {
		res.Users = append(res.Users, h.modelToResponseUser(user))
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		h.handleError(w, err)
		return
	}
}

func (h *UserHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	requestId := h.extractID(r)
	comm := command.NewUserGetCommand(requestId)
	result, err := h.userApplicationService.Get(comm)
	if err != nil {
		h.handleError(w, err)
		return
	}

	res := h.modelToResponseUser(result.User)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		h.handleError(w, err)
		return
	}
}

func (h *UserHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, err)
		return
	}

	var request PostRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.handleError(w, err)
		return
	}

	comm := command.NewUserRegisterCommand(request.Name)
	result, err := h.userApplicationService.Register(comm)
	if err != nil {
		h.handleError(w, err)
		return
	}

	res := h.modelToResponseUser(result.User)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		h.handleError(w, err)
		return
	}
}

func (h *UserHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, err)
		return
	}

	var request PutRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.handleError(w, err)
		return
	}

	comm := command.NewUserUpdateCommand(request.ID, request.Name)
	err = h.userApplicationService.Update(comm)
	if err != nil {
		h.handleError(w, err)
		return
	}
}

func (h *UserHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	requestId := h.extractID(r)
	comm := command.NewUserDeleteCommand(requestId)
	err := h.userApplicationService.Delete(comm)
	if err != nil {
		h.handleError(w, err)
		return
	}
}

func (h *UserHandlerImpl) handleError(w http.ResponseWriter, err error) {
	log.Println(err)
	// TODO: エラー種別に応じてステータスを帰る
	http.Error(w, "Internal Server Error", 500)
}

func (h *UserHandlerImpl) extractID(r *http.Request) string {
	return strings.TrimPrefix(r.URL.Path, "/users/")
}

func (h *UserHandlerImpl) modelToResponseUser(user *model.User) UserResponse {
	id := user.ID()
	name := user.Name()
	userType := user.UserType()
	return UserResponse{
		ID:       id.Value(),
		Name:     name.Value(),
		UserType: userType.Value(),
	}
}

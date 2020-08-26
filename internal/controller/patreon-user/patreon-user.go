package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	patreon_user_service "patreon-statistics/internal/service/patreon-user"
)

type PatreonUserController interface {
	GetOne(w http.ResponseWriter, r *http.Request)
}

type patreonUserController struct {
	patreonUserService patreon_user_service.PatreonUserService
}

func (c *patreonUserController) GetOne(w http.ResponseWriter, r *http.Request) {
	println(111)
	userId, ok := mux.Vars(r)["userId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("userId param is required"))
		return
	}

	user, err := c.patreonUserService.GetOne(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(userJson)
}

func NewPatreonUserController(s patreon_user_service.PatreonUserService) PatreonUserController {
	return &patreonUserController{
		patreonUserService: s,
	}
}

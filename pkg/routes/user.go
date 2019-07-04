package routes

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/tomerBZ/web/pkg/interfaces"
	"github.com/tomerBZ/web/pkg/logs"
	"github.com/tomerBZ/web/pkg/utils"
	"net/http"
)

type userRouter struct {
	userService interfaces.UserService
}

func userRouting(u interfaces.UserService) http.Handler {
	userRouter := userRouter{u}
	router := chi.NewRouter()
	router.Put("/", userRouter.createUserHandler)
	router.Get("/{username}", userRouter.getUserHandler)
	router.Put("/{username}", userRouter.updateUserHandler)
	router.Get("/", userRouter.getUsersHandler)

	return router
}

func (ur *userRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ur.userService.CreateUser(&user)
	if err != nil {
		logs.Error.Println(err)
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Json(w, http.StatusOK, err)
}

func (ur *userRouter) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	userNew, errDecode := decodeUser(r)
	if errDecode != nil {
		logs.Error.Println(errDecode)
		utils.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	username := chi.URLParam(r, "username")
	user, err := ur.userService.GetByUsername(username)

	if err != nil {
		logs.Error.Println(err)
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = ur.userService.UpdateUser(user.Id, &userNew)

	utils.Json(w, http.StatusOK, err)
}

func (ur *userRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := ur.userService.GetByUsername(username)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	utils.Json(w, http.StatusOK, user)
}

func (ur *userRouter) getUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := ur.userService.GetUsers()
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	utils.Json(w, http.StatusOK, users)
}

func decodeUser(ur *http.Request) (interfaces.User, error) {
	var u interfaces.User
	if ur.Body == nil {
		return u, errors.New("no request body")
	}
	decoder := json.NewDecoder(ur.Body)
	err := decoder.Decode(&u)

	return u, err
}

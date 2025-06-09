package http

import (
	"encoding/json"
	"net/http"

	"github.com/Wandering-Digital/anthropos/criteria"
	"github.com/Wandering-Digital/anthropos/domain"
	"github.com/Wandering-Digital/anthropos/internal/customerror"
	"github.com/Wandering-Digital/anthropos/internal/response"
	"github.com/Wandering-Digital/anthropos/request"

	"github.com/go-chi/chi"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(r *chi.Mux, userUseCase domain.UserUseCase) {
	handler := &UserHandler{
		UserUseCase: userUseCase,
	}

	r.Route("/v1/users", func(r chi.Router) {
		r.Post("/", handler.create)
	})
}

func (uh *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqCreateUser := request.CreateUser{}
	if err := json.NewDecoder(r.Body).Decode(&reqCreateUser); err != nil {
		_ = response.WithError(w, customerror.NewError(http.StatusBadRequest, "Invalid payload", err))

		return
	}

	if err := reqCreateUser.Validate(); err != nil {
		_ = response.WithError(w, customerror.NewErrors(http.StatusBadRequest, "Invalid payload", err))

		return
	}

	user, err := uh.UserUseCase.Create(ctx, &criteria.CreateUser{})
	if err != nil {
		_ = response.WithError(w, customerror.NewCustomError(err))

		return
	}

	_ = response.WithData(w, http.StatusCreated, response.Response{
		Data: user,
	})
}

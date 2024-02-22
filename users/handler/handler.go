package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	customError "main/errorLogger"
	bloc "main/users/Bloc"
	"main/users/model"
	"net/http"
)

type Handler struct {
	bloc *bloc.Bloc
}

func NewHandler(bloc *bloc.Bloc) *Handler {
	return &Handler{bloc: bloc}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := h.bloc.Get(r.Context(), &model.GetUserRequest{ID: id})
	if err != nil {
		w.WriteHeader(err.ErrorCode)
		w.Write([]byte(customError.MarshalCustomError(err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	errs := json.NewDecoder(r.Body).Decode(&newUser)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.bloc.Create(r.Context(), &model.CreateUserRequest{
		Username: newUser.Username,
		Password: newUser.Password,
		Email:    newUser.Email,
	})
	if err != nil {
		w.WriteHeader(err.ErrorCode)
		w.Write([]byte(customError.MarshalCustomError(err)))
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedUser model.User
	errs := json.NewDecoder(r.Body).Decode(&updatedUser)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.bloc.Update(r.Context(), &model.UpdateUserRequest{
		ID:       id,
		Username: updatedUser.Username,
		Password: updatedUser.Password,
		Email:    updatedUser.Email,
	})
	if err != nil {
		w.WriteHeader(err.ErrorCode)
		w.Write([]byte(customError.MarshalCustomError(err)))
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.bloc.Delete(r.Context(), &model.DeleteUserRequest{ID: id})
	if err != nil {
		w.WriteHeader(err.ErrorCode)
		w.Write([]byte(customError.MarshalCustomError(err)))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func (h *Handler) Upsert(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	errs := json.NewDecoder(r.Body).Decode(&newUser)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.bloc.Upsert(r.Context(), &model.UpsertUserRequest{
		ID:       newUser.ID,
		Username: newUser.Username,
		Password: newUser.Password,
		Email:    newUser.Email,
	})
	if err != nil {
		w.WriteHeader(err.ErrorCode)
		w.Write([]byte(customError.MarshalCustomError(err)))
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	customError "main/errorLogger"
	bloc "main/post/Bloc"
	"main/post/model"
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
	res, err := h.bloc.Get(r.Context(), &model.GetPostRequest{Id: id})
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
	var newPost model.Post
	errs := json.NewDecoder(r.Body).Decode(&newPost)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.bloc.Create(r.Context(), &model.CreatePostRequest{
		Content: newPost.Content,
		Parent:  newPost.Parent,
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
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var newPost model.Post
	errs := json.NewDecoder(r.Body).Decode(&newPost)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.bloc.Update(r.Context(), &model.UpdatePostRequest{
		Id:      newPost.Id,
		Content: newPost.Content,
		Parent:  newPost.Parent,
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
	return
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.bloc.Delete(r.Context(), &model.DeletePostRequest{Id: id})
	if err != nil {
		w.WriteHeader(err.ErrorCode)
		w.Write([]byte(customError.MarshalCustomError(err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (h *Handler) DeleteAllPostOfUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	err := h.bloc.DeleteAllPostOfUser(r.Context(), &model.DeleteAllPostOfUserRequest{UserId: userId})
	if err != nil {
		w.WriteHeader(err.ErrorCode)
		w.Write([]byte(customError.MarshalCustomError(err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

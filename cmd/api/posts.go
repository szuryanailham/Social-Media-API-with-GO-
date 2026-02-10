package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/szuryanailham/social/internal/env/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}



func (app *Application) CreatePostHandler(w http.ResponseWriter, r *http.Request){
	var payload CreatePostPayload
	if err := readJSON(w, r,&payload) ; err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
	app.badRequestResponse(w,r,err)
	return
	}

	post := &store.Post{
		Title:payload.Title,
		Content: payload.Content, 
		Tags : payload.Tags,
		UsersID:1,
}
	ctx := r.Context()

	if err := app.store.Posts.Create(ctx,post); err != nil {
	app.internalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *Application)GetPostHandler(w http.ResponseWriter, r *http.Request) {
idParams := chi.URLParam(r,"postID")
id , err := strconv.ParseInt(idParams, 10, 64)

if err != nil {
	app.internalServerError(w, r, err)
}
ctx := r.Context()

post , err  := app.store.Posts.GetByID(ctx, id)

if err != nil {
	switch {
	case errors.Is(err, store.ErrNotFound):
		app.badRequestResponse(w, r, err)
	default:
		app.notFoundResponse(w,r , err)
	}
	return
}

comments, err := app.store.Comments.GetByPostID(ctx, id)
if err != nil {
	app.internalServerError(w, r, err)
}

post.Comments = comments

if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *Application) deletePostHandler(w http.ResponseWriter, r*http.Request) {
	idParams := chi.URLParam(r,"postID")
	id , err := strconv.ParseInt(idParams, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
	}
	ctx := r.Context()
	if err := app.store.Posts.Delete(ctx, id); err != nil {
		switch{
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w,r, err)
		default:
			app.internalServerError(w,r,err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}


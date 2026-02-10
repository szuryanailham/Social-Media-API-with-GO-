package main

import (
	"log"
	"net/http"
)

func (app *Application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server Error : %s path:%s error : %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError,"the server encountered a problem")
}

func (app *Application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad request error : %s path:%s error : %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest,err.Error())
}

func (app *Application)notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not found error : %s path:%s error : %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusNotFound,"not found")
}
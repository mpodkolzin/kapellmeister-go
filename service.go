package main

import (
	"context"
	"encoding/json"
	"fmt"
	"kapellmeister-go/model"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

var jobs = make(map[int]*model.Job)

type Service struct {
	Srv    *http.Server
	router *mux.Router
}

func (s *Service) Run() {
	http.ListenAndServe(":9090", s.router)
}

func (s *Service) Stop(ctx context.Context) {
	s.Srv.Shutdown(ctx)
}

func NewService() (Service, error) {

	s := Service{}
	s.router = mux.NewRouter()
	s.Srv = &http.Server{Handler: s.router}
	s.initializeRoutes()
	return s, nil
}

func (s *Service) initializeRoutes() {
	s.router.HandleFunc("/ping", pingHandler).Methods("GET")
	s.router.HandleFunc("/jobs/{id:[0-9]+}", jobsHandler).Methods("POST")
	s.router.HandleFunc("/jobs/{id:[0-9]+}", jobsGetHandler).Methods("GET")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi there 2"))
}

func jobsHandler(w http.ResponseWriter, r *http.Request) {
	var j model.Job
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&j); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	jobs[id] = &j

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hi there %s", j.ID)))
}

func jobsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	res, _ := json.Marshal(jobs[id])
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

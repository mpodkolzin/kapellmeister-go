package main

import (
	"encoding/json"
	"fmt"
	"kapellmeister/model"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

var jobs = make(map[int]*model.Job)

type Service struct {
	Router *mux.Router
}

func (s *Service) Run() {
	http.ListenAndServe(":9090", s.Router)
}

func NewService() (Service, error) {

	s := Service{}
	s.Router = mux.NewRouter()
	s.initializeRoutes()
	return s, nil
}

func (a *Service) initializeRoutes() {
	a.Router.HandleFunc("/ping", pingHandler).Methods("POST")
	a.Router.HandleFunc("/jobs/{id:[0-9]+}", jobsHandler).Methods("POST")
	a.Router.HandleFunc("/jobs/{id:[0-9]+}", jobsGetHandler).Methods("GET")
	//a.Router.HandleFunc("/users/{id:[0-9]+}", jobsHandler).Methods("POST")
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
	w.Write([]byte(fmt.Sprintf("Hi there %s", j.StartTime)))
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

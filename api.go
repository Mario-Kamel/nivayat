package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleGetAccounts)).Methods("GET")
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleCreateAccount)).Methods("POST")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleDeleteAccount)).Methods("DELETE")

	log.Println("Running on port ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccounts(w, r)
	} else if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %v", r.Method)
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}
	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createAccountRequest); err != nil {
		return err
	}

	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName, createAccountRequest.Address)
	account, err := s.store.CreateAccount(account)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusCreated, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	if err = s.store.DeleteAccount(id); err != nil {

	}
	return WriteJson(w, http.StatusNoContent, map[string]int{"deleted": id})
}

func getId(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1, fmt.Errorf("couldn't parse id %v", idStr)
	}
	return id, nil
}

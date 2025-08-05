package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunctions func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunctions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// Handle the error by writing a JSON response
			apiError := APIError{Error: err.Error()}
			if writeErr := writeJSON(w, http.StatusInternalServerError, apiError); writeErr != nil {
				http.Error(w, writeErr.Error(), http.StatusInternalServerError)
			}
		}
	}
}

type APIError struct {
	Error string
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	log.Println("Initializing API server...")

	// Define routes
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	log.Println("Starting server on", s.listenAddr)

	// Start the server
	http.ListenAndServe(s.listenAddr, router)
}

// handleAccount routes requests to the appropriate handler based on the HTTP method
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetAccount(w, r)
	}
	if r.Method == http.MethodPost {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}
	if r.Method == http.MethodPut {
		return s.handleTransfer(w, r)
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return nil
}

// Handlers for account operations
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	var account = newAccount("John", "Doe")
	log.Println("Get account details for:", account.FirstName, account.LastName)
	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

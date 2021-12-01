package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MikolajLuczko/go-rest-api/internal/transaction"
	"github.com/gorilla/mux"
)

// stores pointers to the register service
type Handler struct {
	Router  *mux.Router
	Service *transaction.Service
}

// stores responses from the API
type Response struct {
	Message string
	Error   string
}

// return pointer to a new handler
func NewHandler(service *transaction.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// sets up routes for the app
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "All good"}); err != nil {
			panic(err)
		}
	})
	h.Router.HandleFunc("/api/transaction", h.GetAllTransactions).Methods("GET")
	h.Router.HandleFunc("/api/transaction/{id}", h.GetTransaction).Methods("GET")
	h.Router.HandleFunc("/api/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
	h.Router.HandleFunc("/api/transaction/{id}", h.UpdateTransaction).Methods("PUT")
	h.Router.HandleFunc("/api/transaction", h.PostTransaction).Methods("POST")
}

// retrieve transaction from db by id
func (h *Handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing the transaction ID", err)
		return
	}
	transaction, err := h.Service.GetTransaction(uint(id))
	if err != nil {
		sendErrorResponse(w, "Error retrieving transaction by ID", err)
		return
	}
	if err := sendOkResponse(w, transaction); err != nil {
		panic(err)
	}
}

// retrieves all transactions from the db
func (h *Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.Service.GetAllTransactions()
	if err != nil {
		sendErrorResponse(w, "Error retrieving the transactions", err)
		return
	}
	if err := sendOkResponse(w, transactions); err != nil {
		panic(err)
	}
}

// posts a transaction to the db
func (h *Handler) PostTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction transaction.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		sendErrorResponse(w, "Error decoding the request body", err)
		return
	}

	transaction, err := h.Service.PostTransaction(transaction)
	if err != nil {
		sendErrorResponse(w, "Error posting new transaction", err)
		return
	}
	if err := sendOkResponse(w, transaction); err != nil {
		panic(err)
	}
}

// updates transaction by id
func (h *Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction transaction.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		sendErrorResponse(w, "Error decoding the request body", err)
		return
	}

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing the transaction ID", err)
		return
	}
	transaction, err = h.Service.UpdateTransaction(uint(id), transaction)
	if err != nil {
		sendErrorResponse(w, "Error updating transaction", err)
		return
	}
	if err := sendOkResponse(w, transaction); err != nil {
		panic(err)
	}
}

// deletes transaction from db by id
func (h *Handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing the transaction ID", err)
		return
	}
	err = h.Service.DeleteTransaction(uint(id))
	if err != nil {
		sendErrorResponse(w, "Error deleting transaction", err)
		return
	}

	if err := sendOkResponse(w, Response{Message: "Transaction deleted successfully"}); err != nil {
		panic(err)
	}
}

// sends an OK http response to the caller
func sendOkResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

// sends an error http response to the caller
func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}

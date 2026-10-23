package http_server

import (
	"docker/Company"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	company *Company.Company
}

func NewHandler(company *Company.Company) *Handler {
	return &Handler{company}
}

func (h *Handler) AddHTTP(w http.ResponseWriter, r *http.Request) {
	var worker Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		ErrorHTTPServer(w, err, http.StatusBadRequest)
		return
	}
	h.company.Add(worker.FullName, worker.Position)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "worker added"}); err != nil {
		ErrorHTTPServer(w, err, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetWorkersHTTP(w http.ResponseWriter, r *http.Request) {
	workers := h.company.GetWorkers()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(workers); err != nil {
		ErrorHTTPServer(w, err, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) RemoveHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ErrorHTTPServer(w, err, http.StatusBadRequest)
		return
	}
	h.company.DeleteWorker(idInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "worker removed"}); err != nil {
		ErrorHTTPServer(w, err, http.StatusInternalServerError)
		return
	}
}

func StartHTTPServer(handler http.Handler) error {

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", handler)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func ErrorHTTPServer(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	error_ := ErrorRequest{
		Message:   err.Error(),
		TimeError: time.Now(),
	}
	if err_ := json.NewEncoder(w).Encode(error_); err_ != nil {
		panic(err_)
	}
}

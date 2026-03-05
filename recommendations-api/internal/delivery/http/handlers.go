package httpdelivery

import (
	"encoding/json"
	"errors"
	"net/http"

	"book-recommendation-system/recommendations-api/internal/domain/entities"
	"book-recommendation-system/recommendations-api/internal/usecase"
)

type Handlers struct {
	service *usecase.Service
}

func NewHandlers(service *usecase.Service) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "service": "recommendations-api"})
}

func (h *Handlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}
	created, err := h.service.CreateBook(r.Context(), book)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handlers) ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.ListBooks(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, books)
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}
	created, err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handlers) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, users)
}

func (h *Handlers) TriggerTraining(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.TriggerTraining(r.Context())
	if err != nil {
		if errors.Is(err, usecase.ErrDownstreamFailure) {
			writeJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handlers) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	resp, err := h.service.GetRecommendations(r.Context(), userID)
	if err != nil {
		if err == usecase.ErrUserNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		if errors.Is(err, usecase.ErrDownstreamFailure) {
			writeJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handlers) CreatePurchase(w http.ResponseWriter, r *http.Request) {
	var purchase entities.Purchase
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}
	created, err := h.service.CreatePurchase(r.Context(), purchase)
	if err != nil {
		if err == usecase.ErrUserNotFound || err == usecase.ErrBookNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handlers) ListPurchasesByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	purchases, err := h.service.ListPurchasesByUser(r.Context(), userID)
	if err != nil {
		if err == usecase.ErrUserNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, purchases)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

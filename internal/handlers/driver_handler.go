package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"f1-tracker-backend/internal/models"
	"f1-tracker-backend/internal/repository"

	"github.com/go-chi/chi/v5"
)

type DriverHandler struct {
	repo *repository.DriverRepo
}

func NewDriverHandler(repo *repository.DriverRepo) *DriverHandler {
	return &DriverHandler{repo: repo}
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, models.ErrorResponse{
		Error:   http.StatusText(status),
		Message: msg,
	})
}

func (h *DriverHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))

	filter := repository.DriverFilter{
		Search: q.Get("q"),
		Sort:   q.Get("sort"),
		Order:  strings.ToLower(q.Get("order")),
		Page:   page,
		Limit:  limit,
	}

	drivers, total, err := h.repo.GetAll(filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al obtener drivers")
		return
	}

	if drivers == nil {
		drivers = []models.Driver{}
	}

	totalPages := total / filter.Limit
	if total%filter.Limit != 0 {
		totalPages++
	}

	writeJSON(w, http.StatusOK, models.PaginatedResponse{
		Data:       drivers,
		Page:       filter.Page,
		Limit:      filter.Limit,
		Total:      total,
		TotalPages: totalPages,
	})
}

func (h *DriverHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	driver, err := h.repo.GetByID(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al obtener driver")
		return
	}
	if driver == nil {
		writeError(w, http.StatusNotFound, "Driver no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, driver)
}

func (h *DriverHandler) Create(w http.ResponseWriter, r *http.Request) {
	var d models.Driver
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Validación server-side
	if strings.TrimSpace(d.Name) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'name' es requerido")
		return
	}
	if strings.TrimSpace(d.Team) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'team' es requerido")
		return
	}
	if d.Number <= 0 {
		writeError(w, http.StatusBadRequest, "Campo 'number' debe ser un entero positivo")
		return
	}

	created, err := h.repo.Create(&d)
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			writeError(w, http.StatusConflict, "Número de driver ya existe")
			return
		}
		writeError(w, http.StatusInternalServerError, "Falla al crear driver")
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (h *DriverHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var d models.Driver
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		writeError(w, http.StatusBadRequest, "JSON body inválido")
		return
	}

	if strings.TrimSpace(d.Name) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'name' es requerido")
		return
	}
	if strings.TrimSpace(d.Team) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'team' es requerido")
		return
	}
	if d.Number <= 0 {
		writeError(w, http.StatusBadRequest, "Campo 'number' debe ser un entero positivo")
		return
	}

	updated, err := h.repo.Update(id, &d)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al actualizar driver")
		return
	}
	if updated == nil {
		writeError(w, http.StatusNotFound, "Driver no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (h *DriverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	deleted, err := h.repo.Delete(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al eliminar driver")
		return
	}
	if !deleted {
		writeError(w, http.StatusNotFound, "Driver no encontrado")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

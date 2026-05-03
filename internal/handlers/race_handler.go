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

type RaceHandler struct {
	repo *repository.RaceRepo
}

func NewRaceHandler(repo *repository.RaceRepo) *RaceHandler {
	return &RaceHandler{repo: repo}
}

func (h *RaceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))

	filter := repository.RaceFilter{
		Search: q.Get("q"),
		Sort:   q.Get("sort"),
		Order:  strings.ToLower(q.Get("order")),
		Page:   page,
		Limit:  limit,
	}

	races, total, err := h.repo.GetAll(filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al obtener carreras")
		return
	}

	if races == nil {
		races = []models.Race{}
	}

	totalPages := total / filter.Limit
	if total%filter.Limit != 0 {
		totalPages++
	}

	writeJSON(w, http.StatusOK, models.PaginatedResponse{
		Data:       races,
		Page:       filter.Page,
		Limit:      filter.Limit,
		Total:      total,
		TotalPages: totalPages,
	})
}

func (h *RaceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	race, err := h.repo.GetByID(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch race")
		return
	}
	if race == nil {
		writeError(w, http.StatusNotFound, "Race not found")
		return
	}

	writeJSON(w, http.StatusOK, race)
}

func (h *RaceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var race models.Race
	if err := json.NewDecoder(r.Body).Decode(&race); err != nil {
		writeError(w, http.StatusBadRequest, "JSON body inválido")
		return
	}

	if strings.TrimSpace(race.GrandPrix) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'grand_prix' es requerido")
		return
	}
	if strings.TrimSpace(race.Circuit) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'circuit' es requerido")
		return
	}
	if strings.TrimSpace(race.Country) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'country' es requerido")
		return
	}
	if strings.TrimSpace(race.RaceDate) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'race_date' es requerido")
		return
	}

	created, err := h.repo.Create(&race)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al crear carrera")
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (h *RaceHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID carrera inválido")
		return
	}

	var race models.Race
	if err := json.NewDecoder(r.Body).Decode(&race); err != nil {
		writeError(w, http.StatusBadRequest, "JSON body inválido")
		return
	}

	if strings.TrimSpace(race.GrandPrix) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'grand_prix' es requerido")
		return
	}
	if strings.TrimSpace(race.Circuit) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'circuit' es requerido")
		return
	}
	if strings.TrimSpace(race.Country) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'country' es requerido")
		return
	}
	if strings.TrimSpace(race.RaceDate) == "" {
		writeError(w, http.StatusBadRequest, "Campo 'race_date' es requerido")
		return
	}

	updated, err := h.repo.Update(id, &race)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al actualizar carrera")
		return
	}
	if updated == nil {
		writeError(w, http.StatusNotFound, "Carrera no encontrada")
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (h *RaceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID carrera inválido")
		return
	}

	deleted, err := h.repo.Delete(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al eliminar carrera")
		return
	}
	if !deleted {
		writeError(w, http.StatusNotFound, "Carrera no encontrada")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

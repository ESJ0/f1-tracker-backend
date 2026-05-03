package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"f1-tracker-backend/internal/models"
	"f1-tracker-backend/internal/repository"

	"github.com/go-chi/chi/v5"
)

type ResultHandler struct {
	repo *repository.ResultRepo
}

func NewResultHandler(repo *repository.ResultRepo) *ResultHandler {
	return &ResultHandler{repo: repo}
}

func (h *ResultHandler) GetByRace(w http.ResponseWriter, r *http.Request) {
	raceID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID carrera inválido")
		return
	}

	results, err := h.repo.GetByRace(raceID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al obtener resultados")
		return
	}

	if results == nil {
		results = []models.ResultDetail{}
	}

	writeJSON(w, http.StatusOK, results)
}

func (h *ResultHandler) GetByDriver(w http.ResponseWriter, r *http.Request) {
	driverID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID conductor inválido")
		return
	}

	results, err := h.repo.GetByDriver(driverID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al obtener resultados")
		return
	}

	if results == nil {
		results = []models.ResultDetail{}
	}

	writeJSON(w, http.StatusOK, results)
}

func (h *ResultHandler) Create(w http.ResponseWriter, r *http.Request) {
	var res models.Result
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		writeError(w, http.StatusBadRequest, "JSON body inválido")
		return
	}

	if res.DriverID <= 0 {
		writeError(w, http.StatusBadRequest, "Campo 'driver_id' es requerido")
		return
	}
	if res.RaceID <= 0 {
		writeError(w, http.StatusBadRequest, "Campo 'race_id' es requerido	")
		return
	}
	if res.Position <= 0 || res.Position > 20 {
		writeError(w, http.StatusBadRequest, "Campo 'position' debe estar entre 1 y 20")
		return
	}
	if res.Points < 0 {
		writeError(w, http.StatusBadRequest, "Campo 'points' no puede ser negativo")
		return
	}

	created, err := h.repo.Create(&res)
	if err != nil {
		if err.Error() != "" {
			writeError(w, http.StatusConflict, "Resultado para este conductor y carrera ya existe")
			return
		}
		writeError(w, http.StatusInternalServerError, "Falla al crear resultado")
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (h *ResultHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID resultado inválido")
		return
	}

	deleted, err := h.repo.Delete(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Falla al eliminar resultado")
		return
	}
	if !deleted {
		writeError(w, http.StatusNotFound, "Resultado no encontrado")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

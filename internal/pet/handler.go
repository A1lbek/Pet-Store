package pet

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Pets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		var input Pet
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		created := h.service.CreatePet(input)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)

	case http.MethodGet:
		pets := h.service.GetAllPets()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pets)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) PetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/pets/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pet, ok := h.service.GetPetByID(id)
	if !ok {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
}

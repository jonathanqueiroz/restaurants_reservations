package handlers

import (
	"encoding/json"
	"net/http"
	"restaurants-service/internal/domain/models"
	"restaurants-service/internal/usecases"
	"strconv"
)

type RestaurantHandler struct {
	createRestaurantUseCase  usecases.CreateRestaurantUseCase
	getRestaurantByIDUseCase usecases.GetRestaurantByIDUseCase
	getAllRestaurantsUseCase usecases.GetAllRestaurantsUseCase
}

func NewRestaurantHandler(
	createRestaurantUseCase usecases.CreateRestaurantUseCase,
	getRestaurantByIDUseCase usecases.GetRestaurantByIDUseCase,
	getAllRestaurantsUseCase usecases.GetAllRestaurantsUseCase,
) *RestaurantHandler {
	return &RestaurantHandler{
		createRestaurantUseCase:  createRestaurantUseCase,
		getRestaurantByIDUseCase: getRestaurantByIDUseCase,
		getAllRestaurantsUseCase: getAllRestaurantsUseCase,
	}
}

func (h *RestaurantHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var restaurant models.Restaurant
		if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.createRestaurantUseCase.Execute(&restaurant); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case "GET":
		idParam := r.URL.Query().Get("id")
		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
				return
			}
			restaurant, err := h.getRestaurantByIDUseCase.Execute(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(restaurant); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			restaurants, err := h.getAllRestaurantsUseCase.Execute()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(restaurants); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

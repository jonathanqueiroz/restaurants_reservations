package handlers

import (
	"encoding/json"
	"net/http"
	"reservations-service/internal/domain/models"
	"reservations-service/internal/infrastructure/mq"
	"reservations-service/internal/usecases"
	"strconv"
)

type ReservationHandler struct {
	createReservationUseCase  usecases.CreateReservationUseCase
	getReservationByIDUseCase usecases.GetReservationByIDUseCase
	getAllReservationsUseCase usecases.GetAllReservationsUseCase
}

func NewReservationHandler(
	createReservationUseCase usecases.CreateReservationUseCase,
	getReservationByIDUseCase usecases.GetReservationByIDUseCase,
	getAllReservationsUseCase usecases.GetAllReservationsUseCase,
) *ReservationHandler {
	return &ReservationHandler{
		createReservationUseCase:  createReservationUseCase,
		getReservationByIDUseCase: getReservationByIDUseCase,
		getAllReservationsUseCase: getAllReservationsUseCase,
	}
}

func (h *ReservationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var reservation models.Reservation
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.createReservationUseCase.Execute(&reservation); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Publish message to RabbitMQ
		message, err := json.Marshal(reservation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := mq.PublishMessage("reservation_created", message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case "GET":
		idParam := r.URL.Query().Get("id")
		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				http.Error(w, "Invalid reservation ID", http.StatusBadRequest)
				return
			}
			reservation, err := h.getReservationByIDUseCase.Execute(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(reservation); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			reservations, err := h.getAllReservationsUseCase.Execute()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(reservations); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

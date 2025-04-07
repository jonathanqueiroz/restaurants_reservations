package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"users-service/internal/domain/models"
	"users-service/internal/infrastructure/mq"
	"users-service/internal/usecases"
)

type UserHandler struct {
	createUserUseCase  usecases.CreateUserUseCase
	getUserByIDUseCase usecases.GetUserByIDUseCase
	getAllUsersUseCase usecases.GetAllUsersUseCase
}

func NewUserHandler(
	createUserUseCase usecases.CreateUserUseCase,
	getUserByIDUseCase usecases.GetUserByIDUseCase,
	getAllUsersUseCase usecases.GetAllUsersUseCase,
) *UserHandler {
	return &UserHandler{
		createUserUseCase:  createUserUseCase,
		getUserByIDUseCase: getUserByIDUseCase,
		getAllUsersUseCase: getAllUsersUseCase,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.createUserUseCase.Execute(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Publish message to RabbitMQ
		message, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := mq.PublishMessage("user_created", message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case "GET":
		idParam := r.URL.Query().Get("id")
		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}
			user, err := h.getUserByIDUseCase.Execute(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			users, err := h.getAllUsersUseCase.Execute()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

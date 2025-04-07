package handlers

import (
	"encoding/json"
	"net/http"
	"notifications-service/internal/domain/models"
	"notifications-service/internal/usecases"
	"strconv"
)

type NotificationHandler struct {
	createNotificationUseCase  usecases.CreateNotificationUseCase
	getNotificationByIDUseCase usecases.GetNotificationByIDUseCase
	getAllNotificationsUseCase usecases.GetAllNotificationsUseCase
}

func NewNotificationHandler(
	createNotificationUseCase usecases.CreateNotificationUseCase,
	getNotificationByIDUseCase usecases.GetNotificationByIDUseCase,
	getAllNotificationsUseCase usecases.GetAllNotificationsUseCase,
) *NotificationHandler {
	return &NotificationHandler{
		createNotificationUseCase:  createNotificationUseCase,
		getNotificationByIDUseCase: getNotificationByIDUseCase,
		getAllNotificationsUseCase: getAllNotificationsUseCase,
	}
}

func (h *NotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var notification models.Notification
		if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.createNotificationUseCase.Execute(&notification); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case "GET":
		idParam := r.URL.Query().Get("id")
		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				http.Error(w, "Invalid notification ID", http.StatusBadRequest)
				return
			}
			notification, err := h.getNotificationByIDUseCase.Execute(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(notification); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			notifications, err := h.getAllNotificationsUseCase.Execute()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(notifications); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

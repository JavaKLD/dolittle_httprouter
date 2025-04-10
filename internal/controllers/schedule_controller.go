package controllers

import (
	"dolittle/internal/models"
	"dolittle/internal/services"
	"dolittle/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type ScheduleController struct {
	Service *services.ScheduleService
}

func NewScheduleController(service *services.ScheduleService) *ScheduleController {
	return &ScheduleController{Service: service}
}

func (c *ScheduleController) CreateSchedule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var schedule *models.Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	schedule.Created_at = utils.RoundTime(time.Now())

	id, err := c.Service.CreateSchedule(schedule)
	if err != nil {
		if err.Error() == "Лекарство с таким названием для пользователя уже выписано" {
			http.Error(w, "Запись с таким aid_name для данного пользователя уже существует", http.StatusBadRequest)
			return
		}
		if err.Error() == "Лекарства можно пробивать только с 8:00 до 22:00" {
			http.Error(w, "Заказ можно сделать только с 8 до 22 ", http.StatusInternalServerError)
		}
	}

	w.Header().Set("Contet-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Расписание создано успешно",
	})
}

func (c *ScheduleController) GetUserSchedule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryParam := strings.TrimSpace(r.URL.Query().Get("user_id"))
	if queryParam == "" {
		http.Error(w, `{"error": "Не удалось получить user_id"}`, http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(queryParam, 10, 32)
	if err != nil {
		http.Error(w, `{"error": "Неверный формат user_id"}`, http.StatusBadRequest)
		return
	}

	scheduleID, err := c.Service.FindByUserID(uint(userID))
	if err != nil {
		http.Error(w, `{"error": "Ошибка получения данных"}`, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if len(scheduleID) == 0 {
		json.NewEncoder(w).Encode([]uint{})
		return
	}

	json.NewEncoder(w).Encode(map[string][]uint{"schedules": scheduleID})
}

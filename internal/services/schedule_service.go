package services

import (
	"dolittle/internal/models"
	"dolittle/internal/repos"
	"errors"
	"time"
)

type ScheduleService struct {
	Repo *repos.ScheduleRepo
}

func NewScheduleService(repo *repos.ScheduleRepo) *ScheduleService {
	return &ScheduleService{Repo: repo}
}

func (s *ScheduleService) CreateSchedule(schedule *models.Schedule) (uint, error) {
	exists, err := s.Repo.AidNameExists(schedule.Aid_name, schedule.UserID)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("Лекарство с таким названием для пользователя уже выписано")
	}

	hpur := time.Now().Hour()

	if hpur < 8 || hpur > 22 {
		return 0, errors.New("Лекарства можно пробивать только с 8:00 до 22:00")
	}

	id, err := s.Repo.CreateSchedule(schedule)
	if err != nil {
		return 0, err
	}

	return id, nil
}

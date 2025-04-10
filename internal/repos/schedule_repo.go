package repos

import (
	"database/sql"
	"dolittle/internal/models"
	"fmt"
)

type ScheduleRepo struct {
	db *sql.DB
}

func NewScheduleRepo(db *sql.DB) *ScheduleRepo {
	return &ScheduleRepo{db: db}
}

func (r *ScheduleRepo) CreateSchedule(schedule *models.Schedule) (uint, error) {
	insertQuery := `
		INSERT INTO schedules (aid_name, aid_per_day, user_id, duration)
		VALUES (?, ?, ?, ?)
	`

	res, err := r.db.Exec(insertQuery, schedule.Aid_name, schedule.Aid_per_day, schedule.UserID, schedule.Duration)
	if err != nil {
		return 0, err
	}

	insertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(insertID), nil
}

func (r *ScheduleRepo) AidNameExists(aidName string, userID uint) (bool, error) {
	var count int
	qr := `
	SELECT COUNT(*)
	FROM schedules
	WHERE aid_name = ? AND user_id = ? 
	`

	err := r.db.QueryRow(qr, aidName, userID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Ошибка при проверке существования записи: %v", err)
	}

	return count > 0, nil
}

func (r *ScheduleRepo) FindByUserID(userID uint) ([]uint, error) {
	var scheduleID []uint
	qr := `
		SELECT id 
		FROM schedules
		WHERE user_id = ?
	`

	rows, err := r.db.Query(qr, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		scheduleID = append(scheduleID, id)
	}

	return scheduleID, nil
}

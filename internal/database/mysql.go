package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла ", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к бд", err)
	}

	CreateScheduleTable(db)

	return db, err
}

func CreateScheduleTable(db *sql.DB) {
	createQuery := `CREATE TABLE IF NOT EXISTS schedules (
				id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
				aid_name LONGTEXT NOT NULL,
				aid_per_day BIGINT NOT NULL,
				user_id BIGINT UNSIGNED NOT NULL,
				duration BIGINT,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
				)`

	_, err := db.Exec(createQuery)
	if err != nil {
		log.Fatal("Ошибка создания таблицы schedules", err)
	}
}

package main

import (
	"dolittle/internal/controllers"
	"dolittle/internal/database"
	"dolittle/internal/repos"
	"dolittle/internal/services"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка инициализации бд", err)
	}

	defer db.Close()

	repo := repos.NewScheduleRepo(db)
	service := services.NewScheduleService(repo)
	controller := controllers.NewScheduleController(service)

	router := httprouter.New()

	router.POST("/schedule", controller.CreateSchedule)

	log.Fatal(http.ListenAndServe(":8080", router))
}

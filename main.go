package main

import (
	"fmt"
	"log"
	"net/http"
	"todogoapp/controllers"
	"todogoapp/models"

	"github.com/gernest/utron"
)

func main() {
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	app.Model.Register(&models.Todo{})
	app.Model.AutoMigrateAll()
	app.AddController(controllers.NewTodo)

	port := fmt.Sprintf(":%d", app.Config.Port)
	log.Fatal(http.ListenAndServe(port, app))

}

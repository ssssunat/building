package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/ssssunat/pkg/handler"
	"github.com/ssssunat/pkg/repository"
	"github.com/ssssunat/pkg/service"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "mysecretpassword",
	})
	if err != nil {
		log.Fatal(err)
	}

	repo := service.NewBuildingRepository(db)
	service := service.NewService(repo)

	handler := handler.NewHandler(service)
	http.ListenAndServe(":8080", handler.InitRoutes())
}

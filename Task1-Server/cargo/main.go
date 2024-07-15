package main

import (
	"cargo/domain"
	"log"
)

func main() {
	dbURI := "mongodb://localhost:27017/"

	repo, err := domain.NewRepo(dbURI)
	if err != nil {
		log.Fatal(err)
	}

	var (
		svc     = domain.NewService(repo)
		handler = domain.NewHandler(svc)
		router  = handler.InitRoutes()
	)

	if err := router.Start(":5002"); err != nil {
		log.Fatal(err)
	}
}

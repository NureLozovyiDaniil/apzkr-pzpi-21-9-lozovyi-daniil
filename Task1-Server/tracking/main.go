package main

import "tracking/domain"

func main() {
	dbURI := "mongodb://localhost:27017/"

	svc, err := domain.NewService(dbURI)
	if err != nil {
		panic(err)
	}

	handler := domain.NewHandler(svc)
	router := handler.InitRoutes()

	if err := router.Start(":5001"); err != nil {
		panic(err)
	}
}

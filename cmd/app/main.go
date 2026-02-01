package main

import (
	"log"
	"net/http"
	"pet_store/internal/pet"
)

func main() {
	repo := pet.NewRepository()
	service := pet.NewService(repo)
	handler := pet.NewHandler(service)

	http.HandleFunc("/pets", handler.Pets)
	http.HandleFunc("/pets/", handler.PetByID)

	log.Println("Pet Store server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

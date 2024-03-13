package main

import (
	"example/hello/api-todo-list/configs"
	"example/hello/api-todo-list/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func mainn() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}

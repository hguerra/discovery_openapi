// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redocly Museum API
 *
 * An imaginary, but delightful Museum API for interacting with museum services and information. Built with love by Redocly.
 *
 * API version: 1.0.0
 * Contact: team@redocly.com
 */

package main

import (
	"log"
	"net/http"

	api "github.com/hguerra/museum/internal/infra/api"
)

func main() {
	log.Printf("Server started")

	EventsAPIService := api.NewEventsAPIService()
	EventsAPIController := api.NewEventsAPIController(EventsAPIService)

	OperationsAPIService := api.NewOperationsAPIService()
	OperationsAPIController := api.NewOperationsAPIController(OperationsAPIService)

	TicketsAPIService := api.NewTicketsAPIService()
	TicketsAPIController := api.NewTicketsAPIController(TicketsAPIService)

	router := api.NewRouter(EventsAPIController, OperationsAPIController, TicketsAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"go-webapi/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetPersonsHandler(t *testing.T) {
	// Router Setup
	router := mux.NewRouter()
	router.HandleFunc("/persons", handlers.GetPersons).Methods("GET")

	tests := []struct {
		name          string
		setup         func()
		expectedCode  int
		expectedError string
	}{
		{
			name: "Success",
			setup: func() {
				// Setup für erfolgreichen Fall
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "DatabaseError",
			setup: func() {
				// Mock DB Fehler simulieren
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: "database error",
		},
		{
			name: "InvalidParameter",
			setup: func() {
				// Ungültige Parameter simulieren
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: "invalid parameters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			req := httptest.NewRequest("GET", "/persons", nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			if tt.expectedError != "" {
				assert.Contains(t, rr.Body.String(), tt.expectedError)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// Test Setup
	// DB Mock initialisieren
	// Andere notwendige Initialisierungen

	m.Run()
}

package models

import "time"

type Driver struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Team        string    `json:"team"`
	Nationality string    `json:"nationality"`
	Number      int       `json:"number"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Race struct {
	ID        int       `json:"id"`
	GrandPrix string    `json:"grand_prix"`
	Circuit   string    `json:"circuit"`
	Country   string    `json:"country"`
	RaceDate  string    `json:"race_date"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Result struct {
	ID         int       `json:"id"`
	DriverID   int       `json:"driver_id"`
	RaceID     int       `json:"race_id"`
	Position   int       `json:"position"`
	Points     int       `json:"points"`
	FastestLap bool      `json:"fastest_lap"`
	CreatedAt  time.Time `json:"created_at"`
}

// Para respuestas de error consistentes
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Para paginación
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
}

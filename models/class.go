package models

import (
	"encoding/json"
	"strings"
	"time"
)

// Class represents the structure of a class
// @Description Class information
// @Tags Class
// @Accept json
// @Produce json
// @Property ID string "ID of the class"
// @Property Name string "Name of the class"
// @Property StartDate string "Start date of the class"
// @Property EndDate string "End date of the class"
// @Property Capacity int "Maximum capacity of the class"
// @Router /classes [get]
type Class struct {
	ID        string `json:"id"`         // @Description ID of the class
	Name      string `json:"name"`       // @Description Name of the class
	StartDate Date   `json:"start_date"` // @Description Start date of the class in YYYY-MM-DD format
	EndDate   Date   `json:"end_date"`   // @Description End date of the class in YYYY-MM-DD format
	Capacity  int    `json:"capacity"`   // @Description Maximum capacity of the class
}

// CreateClassRequest represents the request body for creating a new class
// @Description Class creation request
// @Tags Class
// @Accept json
// @Produce json
// @Property Name string "Name of the class"
// @Property StartDate string "Start date of the class"
// @Property EndDate string "End date of the class"
// @Property Capacity int "Maximum capacity of the class"
type CreateClassRequest struct {
	Name      string `json:"name"`       // @Description Name of the class
	StartDate Date   `json:"start_date"` // @Description Start date of the class in YYYY-MM-DD format
	EndDate   Date   `json:"end_date"`   // @Description End date of the class in YYYY-MM-DD format
	Capacity  int    `json:"capacity"`   // @Description Maximum capacity of the class
}

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) error {
	dateStr := strings.Trim(string(data), "\"")
	parsedTime, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return err
	}
	*d = Date(parsedTime)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(time.DateOnly))
}

func (d Date) Time() time.Time {
	return time.Time(d)
}

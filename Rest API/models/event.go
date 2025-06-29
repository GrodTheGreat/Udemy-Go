package models

import (
	"rest/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (event *Event) Save() error {
	query := `
INSERT INTO events(name, description, location, datetime, user_id)
VALUES(?, ?, ?, ?, ?)
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	event.ID = id

	return nil
}

func GetAllEvents() []Event {
	return events
}
func New() *Event {
	return &Event{}
}

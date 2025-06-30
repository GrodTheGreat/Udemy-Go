package models

import (
	"rest/db"
	"rest/utils"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"datetime"`
	UserID      int64     `json:"userId"`
}

func (event *Event) Save() error {
	query := `
INSERT INTO event(name, description, location, datetime, user_id)
VALUES(?, ?, ?, ?, ?)
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer utils.CheckClose(stmt, &err)

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

func GetAllEvents() ([]Event, error) {
	query := `
SELECT id, name, description, location, datetime, user_id
FROM event
`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer utils.CheckClose(rows, &err)

	var events []Event
	var rawDateTime string

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &rawDateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		event.DateTime, err = time.Parse(time.RFC3339, rawDateTime)

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
SELECT id, name, description, location, datetime, user_id
FROM event
WHERE id = ?
`
	row := db.DB.QueryRow(query, id)

	var event Event
	var rawDateTime string
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &rawDateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	event.DateTime, err = time.Parse(time.RFC3339, rawDateTime)

	return &event, nil
}

func (event *Event) Update() error {
	query := `
UPDATE event
SET name = ?, description = ?, location = ?, datetime = ?, user_id = ?
WHERE id = ?
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer utils.CheckClose(stmt, &err)

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (event *Event) Delete() error {
	query := `
DELETE FROM event
WHERE id = ?
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer utils.CheckClose(stmt, &err)

	_, err = stmt.Exec(event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (event *Event) Register(userId int64) error {
	query := `INSERT INTO registration (event_id, user_id)
VALUES (?, ?)
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer utils.CheckClose(stmt, &err)

	_, err = stmt.Exec(event.ID, userId)

	return err
}

func New(name, description, location string, datetime time.Time, userID int64) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		DateTime:    datetime,
		UserID:      userID,
	}
}

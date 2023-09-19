package dbrepo

import (
	"context"
	"time"

	"github.com/robert-gherlan/go-booking-web-app/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation to database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var newID int

	stmt := `INSERT INTO reservations
		(first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, err
}

// InsertRoomRestriction inserts a room restriction to database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO room_restrictions
		(start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	)

	return err
}

// SearchAvailabilityByDatesByRoomID check room availability based on provided dates and room id
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(roomID int, startDate time.Time, endDate time.Time) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var numRows int
	stmt := `
		SELECT 
			COUNT(id)
		FROM room_restrictions 
		WHERE room_id = $1
			AND $2 < end_date
			AND $3 > start_date
		`

	err := m.DB.QueryRowContext(ctx, stmt, roomID, endDate, startDate).Scan(&numRows)
	if err != nil {
		return false, err
	}

	isAnyAvailability := numRows == 0

	return isAnyAvailability, nil
}

// SearchAvailabilityForAllRoomsByDates check room availabilities based on provided dates
func (m *postgresDBRepo) SearchAvailabilityForAllRoomsByDates(startDate, endDate time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var rooms []models.Room
	stmt := `
		SELECT 
			r.id,
			r.name
		FROM rooms r 
		WHERE r.id NOT IN (SELECT room_id FROM room_reservations rr WHERE $1 < rr.end_date AND $2 > rr.start_date)
		`

	rows, err := m.DB.QueryContext(ctx, stmt, startDate, endDate)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomByID get the room entity by id.
func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room
	query := `SELECT id, room_name, created_at, updated_at FROM rooms WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&room.ID, &room.RoomName, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return room, err
	}

	return room, nil
}

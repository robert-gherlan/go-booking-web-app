package repository

import (
	"time"

	"github.com/robert-gherlan/go-booking-web-app/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(roomID int, startDate time.Time, endDate time.Time) (bool, error)
	SearchAvailabilityForAllRoomsByDates(startDate, endDate time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}

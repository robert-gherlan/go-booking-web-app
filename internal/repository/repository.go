package repository

import (
	"time"

	"github.com/robert-gherlan/go-booking-web-app/internal/models"
)

type DatabaseRepo interface {
	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email, password string) (int, string, error)

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(roomID int, startDate time.Time, endDate time.Time) (bool, error)
	SearchAvailabilityForAllRoomsByDates(startDate, endDate time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)

	AllReservations() ([]models.Reservation, error)
	AllNewReservations() ([]models.Reservation, error)
	GetReservationByID(id int) (models.Reservation, error)
	UpdateReservation(u models.Reservation) error
	DeleteReservationByID(id int) error
	UpdateProcessedForReservation(id, processed int) error
	AllRooms() ([]models.Room, error)
	GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)
	InsertBlockForRoom(id int, startDate time.Time) error
	DeleteBlockByID(id int) error
}

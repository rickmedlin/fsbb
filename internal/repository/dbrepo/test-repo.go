package dbrepo

import (
	"errors"
	"fsbb/internal/models"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// insert a reservation into the database.
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	return 1, nil
}

// insert a room restriction into the database.
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

// returns true if availability exists for room ID, false if no availability
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	return false, nil

}

// SearchAvailabilityForAllRooms returns a slice of available rooms for a given date range.
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room
	return rooms, nil

}

// GetRoomByID searches for a room by ID.
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil

}

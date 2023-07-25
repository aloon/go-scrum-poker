package main

import (
	"testing"
	"time"
)

func TestCleanRooms(t *testing.T) {
	rooms = make(map[string]Room)

	rooms["1"] = Room{
		ID:           "1",
		CreatedAt:    time.Now().Add(-4 * time.Hour),
		UpdatedAt:    time.Now().Add(-4 * time.Hour),
		Participants: []Participant{},
	}
	rooms["2"] = Room{
		ID:           "2",
		CreatedAt:    time.Now().Add(-3 * time.Hour),
		UpdatedAt:    time.Now().Add(-3 * time.Hour),
		Participants: []Participant{},
	}
	rooms["3"] = Room{
		ID:           "3",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Participants: []Participant{},
	}
	rooms["4"] = Room{
		ID:           "4",
		CreatedAt:    time.Now().Add(-5 * time.Hour),
		UpdatedAt:    time.Now().Add(-5 * time.Hour),
		Participants: []Participant{},
	}

	CleanRooms()

	if len(rooms) != 1 || rooms["3"].ID != "3" {
		t.Errorf("cleanning rooms fail")
	}

}

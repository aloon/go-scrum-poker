package main

import (
	"testing"
	"time"
)

func TestShowVotes(t *testing.T) {
	// given
	roomID := "test-room"
	rooms = make(map[string]Room)
	rooms[roomID] = Room{
		ID:        roomID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Participants: []Participant{
			{UserName: "test-user", TempVote: "?"},
			{UserName: "test-user2", TempVote: "5"},
			{UserName: "test-user3", TempVote: "2"},
			{UserName: "test-user4", TempVote: ""},
		},
	}

	// when
	showVotes(roomID)

	// then: the participant's votes should be shown & ordered by vote
	room := rooms[roomID]
	if (room.Participants[0].Vote != "?") || (room.Participants[1].Vote != "2") || (room.Participants[2].Vote != "5") || (room.Participants[3].Vote != "") {
		t.Errorf(
			"Votes were not shown: room.Participants[0].Vote: %s, room.Participants[1].Vote: %s, room.Participants[2].Vote: %s, room.Participants[3].Vote: %s",
			room.Participants[0].Vote,
			room.Participants[1].Vote,
			room.Participants[2].Vote,
			room.Participants[3].Vote,
		)
	}
	if (room.Participants[0].UserName != "test-user") || (room.Participants[1].UserName != "test-user3") || (room.Participants[2].UserName != "test-user2") || (room.Participants[3].UserName != "test-user4") {
		t.Errorf(
			"Votes were not shown: room.Participants[0].UserName: %s, room.Participants[1].UserName: %s, room.Participants[2].UserName: %s, room.Participants[3].UserName: %s",
			room.Participants[0].UserName,
			room.Participants[1].UserName,
			room.Participants[2].UserName,
			room.Participants[3].UserName,
		)
	}
}

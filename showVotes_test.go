package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowVotes(t *testing.T) {
	// given
	roomID := "test-room"
	rooms = make(map[string]Room)
	rooms[roomID] = Room{
		ID:        roomID,
		CreatedAt: time.Now().Add(-2 * time.Hour),
		UpdatedAt: time.Now().Add(-2 * time.Hour),
		Participants: []Participant{
			{UserName: "test-user4", TempVote: ""},
			{UserName: "test-user", TempVote: "?"},
			{UserName: "test-user2", TempVote: "5"},
			{UserName: "test-user3", TempVote: "2"},
		},
	}

	// when
	showVotes(roomID)

	// then: the participant's votes should be shown & ordered by vote
	room := rooms[roomID]
	assert.Equal(t, "?", room.Participants[0].Vote, "Expected vote for test-user to be '?'")
	assert.Equal(t, "2", room.Participants[1].Vote, "Expected vote for test-user3 to be '2'")
	assert.Equal(t, "5", room.Participants[2].Vote, "Expected vote for test-user2 to be '5'")
	assert.Equal(t, "", room.Participants[3].Vote, "Expected vote for test-user4 to be ''")

	assert.Equal(t, "test-user", room.Participants[0].UserName, "Expected participant at index 0 to be 'test-user'")
	assert.Equal(t, "test-user3", room.Participants[1].UserName, "Expected participant at index 1 to be 'test-user3'")
	assert.Equal(t, "test-user2", room.Participants[2].UserName, "Expected participant at index 2 to be 'test-user2'")
	assert.Equal(t, "test-user4", room.Participants[3].UserName, "Expected participant at index 3 to be 'test-user4'")

	assert.True(t, time.Until(room.UpdatedAt) <= time.Minute, "UpdatedAt was not updated as expected")
}

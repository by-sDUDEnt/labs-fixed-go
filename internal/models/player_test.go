package models

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestPlayer(t *testing.T) {
	// Create a test User object
	user := User{
		ID:        uuid.New(),
		Username:  "john_doe",
		CreatedAt: time.Now(),
	}

	// Create a Player object using the User
	player := Player{
		User: user,
	}

	// Test for User field
	if player.User != user {
		t.Errorf("Expected User: %+v, got: %+v", user, player.User)
	}

	// Test for ID field
	if player.ID != user.ID {
		t.Errorf("Expected ID: %d, got: %d", user.ID, player.ID)
	}

	// Test for Username field
	if player.Username != user.Username {
		t.Errorf("Expected Username: %s, got: %s", user.Username, player.Username)
	}

	if player.CreatedAt != user.CreatedAt {
		t.Errorf("Expected Username: %s, got: %s", user.Username, player.Username)
	}

}

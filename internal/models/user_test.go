package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestUser(t *testing.T) {
	// Create a UUID for the ID field
	id := uuid.New()

	// Set up a test User object
	user := User{
		ID:        id,
		Username:  "john_doe",
		CreatedAt: time.Now(),
	}

	// Test for ID field
	if user.ID != id {
		t.Errorf("Expected ID: %s, got: %s", id.String(), user.ID.String())
	}

	// Test for Username field
	expectedUsername := "john_doe"
	if user.Username != expectedUsername {
		t.Errorf("Expected Username: %s, got: %s", expectedUsername, user.Username)
	}

	// Test for CreatedAt field
	if user.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
}

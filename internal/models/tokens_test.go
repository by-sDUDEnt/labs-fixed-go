package models

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestToken_IsValid(t *testing.T) {
	type fields struct {
		Hash          []byte
		UserID        uuid.UUID
		Scope         string
		CreatedAt     time.Time
		LastVisitedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				CreatedAt:     time.Now().Add(-time.Minute),
				LastVisitedAt: time.Now().Add(-time.Minute),
				Scope:         ScopeSessionUser,
			},
			want: true,
		},
		{
			name: "invalid",
			fields: fields{
				CreatedAt:     time.Now().Add(-time.Minute),
				LastVisitedAt: time.Now().Add(-time.Minute),
				Scope:         "invalid",
			},
			want: false,
		},
		{
			name: "invalid",
			fields: fields{
				CreatedAt:     time.Now().Add(-time.Hour),
				LastVisitedAt: time.Now().Add(-time.Hour),
				Scope:         "invalid",
			},
			want: false,
		},
		{
			name: "invalid",
			fields: fields{
				CreatedAt:     time.Now().Add(-time.Minute),
				LastVisitedAt: time.Now().Add(-time.Hour),
				Scope:         "invalid",
			},
			want: false,
		},
		{
			name: "invalid",
			fields: fields{
				CreatedAt:     time.Now().Add(-time.Hour),
				LastVisitedAt: time.Now().Add(-time.Minute),
				Scope:         "invalid",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := &Token{
				Hash:          tt.fields.Hash,
				UserID:        tt.fields.UserID,
				Scope:         tt.fields.Scope,
				CreatedAt:     tt.fields.CreatedAt,
				LastVisitedAt: tt.fields.LastVisitedAt,
			}
			if got := tok.IsValid(); got != tt.want {
				t.Errorf("Token.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

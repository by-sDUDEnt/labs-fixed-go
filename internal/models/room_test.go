package models

import (
	"github.com/google/uuid"
	"testing"
)

func TestRoom_IsFull(t *testing.T) {
	type fields struct {
		ID      uuid.UUID
		Players []uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "room is full",
			fields: fields{
				ID:      uuid.New(),
				Players: []uuid.UUID{uuid.New(), uuid.New()},
			},
			want: true,
		},
		{
			name: "room is not full",
			fields: fields{
				ID:      uuid.New(),
				Players: []uuid.UUID{uuid.New()},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &Room{
				ID:        tt.fields.ID,
				PlayerIDs: tt.fields.Players,
			}

			if got := r.IsFull(); got != tt.want {
				t.Errorf("Room.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_ContainsPlayer(t *testing.T) {
	type fields struct {
		ID      uuid.UUID
		Players []uuid.UUID
	}
	type args struct {
		userID uuid.UUID
	}
	userID := uuid.New()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"room contains player",
			fields{
				ID:      uuid.New(),
				Players: []uuid.UUID{userID},
			},
			args{
				userID: userID,
			},
			true,
		},
		{
			"room does not contain player",
			fields{
				ID:      uuid.New(),
				Players: []uuid.UUID{uuid.New()},
			},
			args{
				userID: uuid.New(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &Room{
				ID:        tt.fields.ID,
				PlayerIDs: tt.fields.Players,
			}

			if got := r.ContainsPlayer(tt.args.userID); got != tt.want {
				t.Errorf("Room.ContainsPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

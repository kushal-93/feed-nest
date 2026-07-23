package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/kushal-93/feed-nest/internal/database"
)

type modelUser struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	APIKey    string    `json:"apiKey"`
}

func databaseUserToUser(user database.User) modelUser {
	return modelUser{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Name:      user.Name,
		APIKey:    user.ApiKey,
	}
}

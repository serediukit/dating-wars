package model

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // "-" means this field won't be included in JSON
	Name string    `json:"name" gorm:"not null"`
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
	Gender    string    `json:"gender" gorm:"not null"`
	Bio       string    `json:"bio"`
	Location  Location  `json:"location" gorm:"embedded"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Location represents user's location
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CityID      uint  `json:"city_id"`
	CountryID   uint  `json:"country_id"`
}

// Photo represents user's photo
type Photo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	URL       string    `json:"url" gorm:"not null"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserPreferences represents user's dating preferences
type UserPreferences struct {
	UserID        uint    `json:"user_id" gorm:"primaryKey"`
	MinAge        uint     `json:"min_age"`
	MaxAge        uint     `json:"max_age"`
	MaxDistance   uint `json:"max_distance"` // in kilometers
	Gender        string  `json:"gender"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
} 
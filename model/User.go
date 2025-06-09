package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type SearchableUser struct {
	ID             int       `json:"id"`
	GenderID       byte      `json:"gender_id"`
	SearchGenderID byte      `json:"search_gender_id"`
	Age            byte      `json:"age"`
	SearchAgeFrom  byte      `json:"search_age_from"`
	SearchAgeTo    byte      `json:"search_age_to"`
	CountryID      byte      `json:"country_id"`
	Latitude       float32   `json:"latitude"`
	Longitude      float32   `json:"longitude"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserData struct {
	User           User           `json:"user"`
	SearchableUser SearchableUser `json:"searchable_user"`
}

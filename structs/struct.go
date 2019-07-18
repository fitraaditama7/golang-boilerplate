package structs

import (
	"time"
)

type Person struct {
	ID        	uint 	`json:"id"gorm:"primary_key"`
	First_Name 	string 	`json:"first_name"`
	Last_Name 	string	`json:"last_name"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	DeletedAt 	*time.Time `json:"deleted_at"`

}

package models

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
	ChangedOn   time.Time `json:"changedon"`
}

func deneme() {
	fmt.Println("sdf")
}

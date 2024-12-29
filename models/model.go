package models

type Employee struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Role string `json:"role"`
}

package model

type User struct {
	ID           int64  `gorm:"primary_key;not_null;auto_increment"`
	Username     string `gorm:"unique_index;not_null"`
	FirstName    string
	HashPassword string
}

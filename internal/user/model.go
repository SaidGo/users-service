package user

import "time"

type User struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"uniqueIndex;size:255;not null"`
	Name      string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

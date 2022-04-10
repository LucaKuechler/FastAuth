package dao

type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)"`
	Password string  `gorm:"not null"`
}

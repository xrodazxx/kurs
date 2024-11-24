package kurs

// Dlya polzovatilya
type User struct {
	Id       int    `json:"_" db:"id"`
	Name     string `json:"name" binding:"reqvired"`
	Username string `json:"username" binding:"reqvired"`
	Password string `json:"password" binding:"reqvired"`
}

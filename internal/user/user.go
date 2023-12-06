package user

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"-"`
}
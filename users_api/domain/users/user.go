package users

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"datecreated"`
}

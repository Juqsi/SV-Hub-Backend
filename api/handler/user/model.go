package user

type User struct {
	Id       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Forename string `json:"forename" db:"forename"`
	Lastname string `json:"lastname" db:"lastname"`
	Telenum  string `json:"telenum" db:"telenum"`
	Salt     string `json:"-" db:"salt"`
	Hash     string `json:"-" db:"hash"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

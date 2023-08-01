package model

type User struct {
	FullName    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	Balance     float64 `json:"balance"`
	Login       string  `json:"login"`
	Password    string  `json:"password"`
}

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

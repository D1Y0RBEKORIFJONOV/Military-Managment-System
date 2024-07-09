package models

type RegisterModel struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirhtDay  string `json:"birht_day"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"-"`
}

type RegisterResponse struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Role         string `json:"role"`
	Password     string `json:"password"`
	BirhtDay     string `json:"birht_day"`
	Email        string `json:"email"`
	Id           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

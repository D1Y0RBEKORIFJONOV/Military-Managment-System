package soldiers_entity

import "time"

type CreateSoldiersReq struct {
	Fname      string    `json:"fname"`
	Lname      string    `json:"l_name"`
	Email      string    `json:"email"`
	Password   string    `json:"pasword"`
	Birthday   time.Time `json:"birthday"`
	Role       string    `json:"role"`
	SecredCode string    `json:"secred_code"`
}
type Soldiers struct {
	ID         string    `json:"id"`
	Fname      string    `json:"fname"`
	Lname      string    `json:"l_name"`
	Email      string    `json:"email"`
	Password   string    `json:"pasword"`
	Birthday   time.Time `json:"birthday"`
	Age        int       `json:"age"`
	Joined_at  time.Time `json:"joined_at"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
	Role       string    `json:"role"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"pasword"`
}

type RegisterReq struct {
	Email      string `json:"email"`
	SecredCode string `json:"secred_code"`
}

type FildValueReq struct {
	Filed string `json:"filed"`
	Value string `json:"value"`
}

type GetAllSoldierRequests struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}

type UpdateSoldierRequests struct {
	ID       string    `json:"id"`
	Fname    string    `json:"fname"`
	Lname    string    `json:"lname"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`
}

type DeleteSoldiersRequest struct {
	ID           string `json:"id"`
	IsHardDelete bool   `json:"is_hard_delete"`
}

package models

type Storehouse struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Price         float32 `json:"price"`
	Amount        int32   `json:"amount"`
	TypeArtillery string  `json:"type_artillery"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	DeletedAt     string  `json:"deleted_at"`
}

type CreateStorehouseReq struct {
	Name          string  `json:"name"`
	Price         float32 `json:"price"`
	Amount        int32   `json:"amount"`
	TypeArtillery string  `json:"type_artillery"`
}

type UpdateStorehouseReq struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Price         float32 `json:"price"`
	Amount        int32   `json:"amount"`
	TypeArtillery string  `json:"type_artillery"`
}

type GetStorehouseReq struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type GetAllStorehouseReq struct {
	Field  string `json:"field"`
	Value  string `json:"value"`
	Offset int64  `json:"page"`
	Limit  int64  `json:"limit"`
}

type GetAllStorehouseRes struct {
	Storehouses []*Storehouse `json:"storehouses"`
	Count       int64         `json:"count"`
}

type DeleteStorehouseReq struct {
	Id string `json:"id"`
}

type Status struct {
	Message string `json:"message"`
}

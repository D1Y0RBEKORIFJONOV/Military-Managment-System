package models

type TopResourceUsageReq struct {
	SortBy    string `json:"sort_by"`
	TopBy     string `json:"top_by"`
	Limit     int64  `json:"limit"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

type TopResourceUsageRes struct {
	SortBy string `json:"sort_by"`
	Count  int64  `json:"count"`
}

type ResourceUsage struct {
	Id         string  `json:"id"`
	SoldierId  string  `json:"soldier_id"`
	StorageId  string  `json:"storage_id"`
	Amount     int64   `json:"amount"`
	TotalPrice float32 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type CreateResourceUsageReq struct {
	SoldierId  string  `json:"soldier_id"`
	StorageId  string  `json:"storage_id"`
	Amount     int64   `json:"amount"`
	TotalPrice float32 `json:"total_price"`
}

type GetAllResourceUsageReq struct {
	Page      int64  `json:"page"`
	Limit     int64  `json:"limit"`
	Field     string `json:"field"`
	Value     string `json:"value"`
	SortBy    string `json:"sort_by"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

type GetAllResourceUsageRes struct {
	ResourceUsages []*ResourceUsage `json:"resource_usages"`
	Count          int64            `json:"count"`
}

type DeleteResourceUsageReq struct {
	Id string `json:"id"`
}

type DeleteResourceUsageRes struct {
	Message string `json:"message"`
}

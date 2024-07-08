package group_entity

type Group struct {
	Id         string `json:"id"`
	SoldiersID string `json:"soldiers_id"`
	GroupName  string `json:"group_name"`
	Size       int64  `json:"size"`
	SizeLimit  int64  `json:"size_limit"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
type CreateGroupRequest struct {
	GroupName string `json:"group_name"`
	SizeLimit int64  `json:"size_limit"`
}

type AddGroupSoldersRequest struct {
	Id       string `json:"id"`
	SolderID string `json:"solder_id"`
}

type DeleteGroupRequest struct {
	Id    string `json:"id"`
	Field string `json:"field"`
	Value string `json:"value"`
}

type GetAllServiceRequest struct {
	Field   string `json:"field"`
	Value   string `json:"value"`
	Page    int64  `json:"page"`
	Limit   int64  `json:"limit"`
	SordBy  string `json:"order_by"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}

type UpdateGroupRequest struct {
	Id        string `json:"id"`
	GroupName string `json:"group_name"`
	SizeLimit int64  `json:"size_limit"`
}

type AddGroupSoldersResponse struct {
	SoldersGroup struct {
		Id       string `json:"id"`
		FnName   string `json:"fn_name"`
		LnName   string `json:"ln_name"`
		Email    string `json:"email"`
		BirthDay string `json:"birth_day"`
		Role     string `json:"role"`
	}
	GroupID string `json:"group_id"`
}

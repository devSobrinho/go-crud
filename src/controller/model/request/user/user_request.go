package request

import requestCommon "github.com/devSobrinho/go-crud/src/controller/model/request/common"

type UserRequest struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserListRequest struct {
	ID              string                        `form:"id" binding:"omitempty,hexadecimal,len=24"`
	Email           string                        `form:"email" binding:"omitempty,email"`
	Size            string                        `form:"size" binding:"omitempty"`
	PaginationQuery requestCommon.PaginationQuery `form:"pagination_query" validate:"omitempty"`
}

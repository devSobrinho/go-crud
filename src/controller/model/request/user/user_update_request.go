package request

type UserUpdateRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	Age  int8   `json:"age" binding:"required,gt=0,lt=130"`
}

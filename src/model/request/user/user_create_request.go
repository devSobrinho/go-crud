package request

type UserCreateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50,containsany=*&^%$#@!"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Age      int8   `json:"age" binding:"required,gt=0,lt=130"`
}

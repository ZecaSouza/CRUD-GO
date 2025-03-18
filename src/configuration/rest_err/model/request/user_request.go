package request

type UserRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6,max=50,containsany=!@#$%^&*"`
    Name     string `json:"name" binding:"required,min=4,max=100"`
    Age      int8   `json:"age" binding:"required,gte=0,lte=120"` 
}

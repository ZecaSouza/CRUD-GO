package routes

import (
	"github.com/ZecaSouza/CRUD-GO/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FinduserEmail)
	r.POST("/user/:userId", controller.CreatdUser)
	r.PUT("/user/:userId", controller.UpdateUserById)
	r.DELETE("/deleterUser/:userId", controller.DeleteUserById)
}
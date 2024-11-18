// Registrasi User

package controller

import (
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"baf-credit-score/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserController struct {
	uc usecase.UserUsecase
	r *gin.RouterGroup
}

func (us *UserController) createHandler(c *gin.Context){
	var payload dto.UserRequestDto
	err := c.ShouldBindJSON(&payload)	
	if err != nil {
		common.SendErrorResponse(c,http.StatusBadRequest,err.Error())
		return 
	}
	if err := us.uc.RegisterUser(payload); err != nil {
		common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
		return 
	}
	common.SendSuccessResponse(c,payload,"Success Register User")
}

func (us *UserController) Route(){
	us.r.POST("/users",us.createHandler)
}

func NewUserController(
	usecase usecase.UserUsecase,
	rg *gin.RouterGroup,
	) *UserController {
	return &UserController{
		uc: usecase,
		r : rg,	
	}
}
// Registrasi User

package controller

import (
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"baf-credit-score/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthController struct {
	uc usecase.AuthenticationUsecase
	r *gin.RouterGroup
}

func (us *AuthController) loginHandler(c *gin.Context){
	var payload dto.LoginDto
	err := c.ShouldBindJSON(&payload)	
	if err != nil {
		common.SendErrorResponse(c,http.StatusBadRequest,err.Error())
		return 
	}
	token, err := us.uc.Login(payload.Email,payload.Password); 
	if err != nil {
		common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
		return 
	}	
	common.SendSuccessResponse(c,token,"Success Login")
}

func (us *AuthController) Route(){
	us.r.POST("/auth/login",us.loginHandler)
}

func NewAuthController(
	usecase usecase.AuthenticationUsecase,
	rg *gin.RouterGroup,
	) *AuthController {
	return &AuthController{
		uc: usecase,
		r : rg,	
	}
}
// Registrasi User

package controller

import (
	"baf-credit-score/delivery/middleware"
	"baf-credit-score/usecase"
	"baf-credit-score/utils/common"
	"baf-credit-score/utils/constant"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc            usecase.UserUsecase
	r             *gin.RouterGroup
	authMiddlware middleware.AuthMiddleware
}

func (us *UserController) listHandler(c *gin.Context) {
	sizeParam := c.DefaultQuery("size", "3")
	pageParam := c.DefaultQuery("page", "1")

	size, err := strconv.Atoi(sizeParam)
	if err != nil || size <= 0 {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid size parameter")
		return
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid page parameter")
		return
	}

	users, paging, err := us.uc.FindAll(size, page)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var customerItems []any
	for _, customer := range users {
		customerItems = append(customerItems, customer)
	}
	common.SendPageResponse(c, customerItems, paging, "Success Get All Users")
}

func (us *UserController) Route() {
	us.r.GET(constant.Users, us.authMiddlware.RequireToken(constant.ADMIN), us.listHandler)
}

func NewUserController(
	usecase usecase.UserUsecase,
	rg *gin.RouterGroup,
	authMiddleware middleware.AuthMiddleware,
) *UserController {
	return &UserController{
		uc:            usecase,
		r:             rg,
		authMiddlware: authMiddleware,
	}
}

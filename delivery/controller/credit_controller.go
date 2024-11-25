package controller

import (
	"baf-credit-score/delivery/middleware"
	"baf-credit-score/model/dto"
	"baf-credit-score/usecase"
	"baf-credit-score/utils/common"
	"baf-credit-score/utils/constant"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreditController struct {
	uc usecase.CreditUsecase
	r             *gin.RouterGroup
	authMiddlware middleware.AuthMiddleware
}

func (cc *CreditController) createHandler(c *gin.Context) {
	var payload dto.CreditRequestDto
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(c,http.StatusBadRequest,err.Error())
        return
	}
	if err := cc.uc.CreateCredit(payload); err != nil {
		common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
	}
	common.SendSuccessResponse(c,payload,"Success Create Credit")
}

func (cc *CreditController) listHandler(c *gin.Context) {
    sizeParam := c.DefaultQuery("size", "3")
    pageParam := c.DefaultQuery("page", "1")

    size, err := strconv.Atoi(sizeParam)
    if err != nil || size <= 0 {        
		common.SendErrorResponse(c,http.StatusBadRequest,"Invalid size parameter")
        return
    }
    
    page, err := strconv.Atoi(pageParam)
    if err != nil || page <= 0 {
        common.SendErrorResponse(c,http.StatusBadRequest,"Invalid page parameter")
        return
    }


    credits, paging,err := cc.uc.FindAll(size,page)
    if err != nil {
        common.SendErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }
    var creditItems []any
	for _, customer := range credits {
		creditItems = append(creditItems, customer)
	}
    common.SendPageResponse(c,creditItems,paging,"Success Get All Credits")
}

func (cc *CreditController) Route() {
	cc.r.POST(constant.CREDIT_PATH, cc.authMiddlware.RequireToken(constant.ADMIN_ROlE), cc.createHandler)
	cc.r.GET(constant.CREDIT_PATH, cc.authMiddlware.RequireToken(constant.ADMIN_ROlE), cc.listHandler)
}

func NewCreditController(
	usecase usecase.CreditUsecase,
	rg *gin.RouterGroup,
	authMiddleware middleware.AuthMiddleware,
) *CreditController {
	return &CreditController{
		uc:            usecase,
		r:             rg,
		authMiddlware: authMiddleware,
	}
}

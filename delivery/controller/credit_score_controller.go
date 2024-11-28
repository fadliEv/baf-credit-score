package controller

import (
	"baf-credit-score/delivery/middleware"
	"baf-credit-score/usecase"
	"baf-credit-score/utils/common"
	"baf-credit-score/utils/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreditScoreController struct {
	uc            usecase.CreditScoreUsecase
	r             *gin.RouterGroup
	authMiddlware middleware.AuthMiddleware
}

func (c *CreditScoreController) getByCustomerHandler(ctx *gin.Context) {
	id := ctx.Param("id")
    credit, err := c.uc.FindScoreByCustomer(id)
    if err != nil {
        common.SendErrorResponse(ctx,http.StatusInternalServerError,err.Error())
        return
    }

    common.SendSuccessResponse(ctx,credit,"Success Get Score Customer")
}


func (cc *CreditScoreController) Route() {
    cc.r.GET(constant.CREDIT_SCORE_ID_PATH,cc.authMiddlware.RequireToken(constant.ADMIN_ROlE), cc.getByCustomerHandler)    
}

func NewCreditScoreController(
    usecase usecase.CreditScoreUsecase,
    r *gin.RouterGroup,
    authMiddleware middleware.AuthMiddleware,
    ) *CreditScoreController {
    return &CreditScoreController{
        uc: usecase,
        r:  r,
        authMiddlware: authMiddleware,
    }
}
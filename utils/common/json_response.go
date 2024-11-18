package common

import (
	"baf-credit-score/utils/payload"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(c *gin.Context, data any, responseType string) {
	c.JSON(http.StatusOK, &payload.SingleResponse{
		Status: payload.Status{
			Code:        http.StatusOK,
			Description: responseType,
		},
		Data: data,
	})
}

func SendPageResponse(c *gin.Context, data []any, paging payload.Paging,responseType string) {
	c.JSON(http.StatusOK, &payload.PagedResponse{
		Status: payload.Status{
			Code:        http.StatusOK,
			Description: responseType,
		},
		Data:   data,
		Paging: paging,
	})
}

func SendErrorResponse(c *gin.Context, code int, errorMessage string) {
	c.AbortWithStatusJSON(code, &payload.Status{
		Code:        code,
		Description: errorMessage,
	})
}

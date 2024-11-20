package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockAuthMiddleware struct {
	mock.Mock
}

func (m *MockAuthMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user", "test-user-id")
		c.Next()
	}
}

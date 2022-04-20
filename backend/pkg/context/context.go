package context

import (
	"context"

	"github.com/gin-gonic/gin"
	"learning/pkg/jwt"
)

type Handler func(c *Context)

type Context struct {
	*gin.Context
	Ctx       context.Context
	UserToken *jwt.UserToken
}

func WrapperHandler(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := jwt.GetToken(c)
		ctx := &Context{
			Ctx:       context.TODO(),
			UserToken: token,
			Context:   c,
		}

		h(ctx)
	}
}

package api

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func (x *API) Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authorization := c.GetHeader("Authorization")
		if authorization == nil {
			c.AbortWithStatusJSON(401, utils.H{
				"code":    0,
				"message": "Please set the token.",
			})
			return
		}

		parts := strings.SplitN(string(authorization), " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(401, utils.H{
				"code":    0,
				"message": "The authorization header is incorrect.",
			})
			return
		}

		if _, err := x.Passport.Verify(parts[1]); err != nil {
			c.AbortWithStatusJSON(401, utils.H{
				"code":    0,
				"message": "The token is inconsistent or has expired.",
			})
			return
		}

		c.Next(ctx)
	}
}

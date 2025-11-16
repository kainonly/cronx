package index

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func (x *Controller) Ping(_ context.Context, c *app.RequestContext) {
	data := M{
		"node":     x.V.Node,
		"hostname": os.Getenv("HOSTNAME"),
		"ts":       time.Now(),
	}
	c.JSON(200, data)
}

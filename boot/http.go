package boot

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/thoohv5/blog/route"
)

func Http() error {
	// gin new
	r := gin.New()

	// gin middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// gin register route
	route.RegisterRoute(r)

	// gin run
	if err := r.Run(":80"); nil != err {
		return fmt.Errorf("http Run err: %w", err)
	}

	return nil
}

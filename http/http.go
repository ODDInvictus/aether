package http

import (
	"fmt"
	"time"

	"github.com/ODDInvictus/aether/logger"
	"github.com/ODDInvictus/aether/spotify"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Init() *gin.Engine {
	r = gin.Default()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s - \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())


	apiRoutes()

	return r
}

func apiRoutes() {
	r.GET("/hello", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "World!",
		})
	})

	r.GET("/state", func (c *gin.Context) {
		state, err := spotify.Current()

		if err != nil {
			c.JSON(500, gin.H{
				"message": fmt.Sprint(err),
			})
		}

		logger.Verbose(fmt.Sprint(state))

		c.JSON(200, gin.H{
			"state": state,
		})
	})

	r.POST("/playlist/play", func (c *gin.Context) {
		var params PlaylistPlay

		if c.ShouldBind(&params) == nil {
			ok, err := spotify.Load(params.SpotifyID, true, true)

			if !ok {
				c.JSON(500, gin.H{
					"message": fmt.Sprint(err),
				})
			}

			c.JSON(200, gin.H{
				"message": "Success",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Invalid playlist id",
			})
		}
	})

	r.POST("/skip", func (c *gin.Context) {
		spotify.Next()
	})
}
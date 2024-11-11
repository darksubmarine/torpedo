package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
)

// HttpServerProvider provides a gin http server
type HttpServerProvider struct {
	app.BaseProvider
	cfg conf.Map

	// server gin instance to be provided
	server *gin.Engine `torpedo.di:"provide"`

	// api router group to add endpoints under /api prefix
	apiV1 *gin.RouterGroup `torpedo.di:"provide,name=APIv1"`

	// binds
	logger log.ILogger `torpedo.di:"bind"`
}

func NewHttpServerProvider(config conf.Map) *HttpServerProvider {
	return &HttpServerProvider{cfg: config}
}

// Provide set the provide instances
func (p *HttpServerProvider) Provide(c app.IContainer) error {
	// gin server default instance
	p.server = gin.Default()

	// Adding a /ping endpoint
	p.server.GET("/ping", func(c *gin.Context) {
		p.logger.Info("/ping has been called")
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// api v1 group
	p.apiV1 = p.server.Group("/api/v1")

	return nil
}

// OnStart launches the gin http server calling the gin.Run method
func (p *HttpServerProvider) OnStart() func() error {
	return func() error {
		go func() {
			port := p.cfg.FetchIntP("port")
			if err := p.server.Run(fmt.Sprintf(":%d", port)); err != nil {
				panic(fmt.Sprintf("error starting HTTP server at %d with error %s", port, err))
			}
		}()

		return nil
	}
}

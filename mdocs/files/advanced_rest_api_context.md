In the context of an HTTP request, **"context"** refers to the additional information, environment, or state that surrounds or supports the request during its lifecycle. It typically includes metadata and parameters that help process the request and manage its execution. 

In summary, the "context" in an HTTP request is any auxiliary information or state passed along with the request to help with its handling, processing, or management across different layers of an application or system.

Torpedo provides an **Execution Context Map** that can be integrated easily with the Gin Gonic Context object and shared with **entity service or even use cases**. 

> This refers to the broader environment in which an HTTP request is being handled. It includes data such as the current user, security roles, server settings, or anything else needed to handle the request within an application.
> 
> For instance, in cloud or microservices architectures, an execution context could include tracing information for distributed systems.

## Adding Execution Context Map to Gin

Torpedo lib provides a "Gin utils" package to do this integration straight forward. The context map is added into the request life cycle
as middleware, for instance:

```go hl_lines="6 47"
package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/http/gin_utils" //(1)!
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

	p.apiV1.Use(gin_utils.WithDataContext()) //(2)!

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

```

1. Import the `gin_utils` package to do the integration.
2. Adding the Execution Data Context on each API v1 request.

## Writing data within Execution Context Map

Once that the context map is set into each request lifecycle it is available to set `key-value` data pairs. For instance, 
following the `sensor` controller:

```go hl_lines="4 6 7"
func (h *InputGin) MyAwesomeENdpoint(c *gin.Context) {
	id := c.Param("id")
	token := c.Request.Header["Token"]
	gin_utils.SetDataContext(c, "token", token) //(1)!

    ctx, _ := gin_utils.GetDataContext(c) //(2)!
    if err := h.srv.Delete(ctx, id); err != nil { //(3)!
        if errors.Is(err, torpedo_lib.ErrIdNotFound) {
            c.JSON(http.StatusNotFound, api.ErrorNotFound(err))
        } else {
            c.JSON(http.StatusInternalServerError, api.ErrorEntityRemove(err))
        }
        return
    }

    c.JSON(http.StatusNoContent, nil)
}
```

1. Setting the token and binding it withing Gin Context, maybe if other middlewares need it. <br>Method definition: `SetDataContext(c *gin.Context, key string, val interface{})`
2. Fetching the Execution Data Context from Gin Context to shared with the service method.
3. Calling service method and passing the Execution Data Context as parameter. In this example, from the service could be possible read the user token.

## Reading data from Execution Context Map

The execution context map is a `sync.Map` struct ready to write and read. The map implements an interface to cast key-value data:

```go
// IDataMap interface to defines DataMap methods.
type IDataMap interface {
	Set(key string, val interface{})
	Get(key string) interface{}
	GetOrElse(key string, val interface{}) interface{}
	GetStringOrElse(key string, val string) string
	GetInt64OrElse(key string, val int64) int64
	GetIntOrElse(key string, val int) int
	GetFloat64OrElse(key string, val float64) float64
	GetBoolOrElse(key string, val bool) bool
}
```

So, to fetch data only needs to know the key which is a string key. Also, the map provides methods to return default values in case that the given key doesn't exist.

Following the previous example we can fetch the token like:

```go hl_lines="4 9 18"
// Delete removes the entity given its id
func (s *ServiceBase) Delete(ctx context.IDataMap, id string) error {
	hook := s.hookDelete()
	hookErr := s.execHookById(hook.beforeDelete, ctx, id) //(1)!
	if hookErr != nil {
		s.logger.Error("before delete hook", "error", hookErr)
	}

	token := ctx.GetStringOrElse("token", "unknown") //(2)!
	
	s.logger.Debug("deleting from repo", "id", id, "user-token", token)
	err := s.repo.DeleteByID(id)
	if err != nil {
		s.logger.Error("deleting from repo", "error", err, "id", id)
		return err // prevents call the after delete hook due to an error
	}

	hookErr = s.execHookById(hook.afterDelete, ctx, id) //(3)!
	if hookErr != nil {
		s.logger.Error("before delete hook", "error", hookErr)
	}

	return err
}
```

1. At basic CRUD operations the context is passed throw the hook functions.
2. Reading the token value or `unknown` value if it is not set. This is only illustrative, not included at basic CRUD operations.
3. At basic CRUD operations the context is passed throw the hook functions.
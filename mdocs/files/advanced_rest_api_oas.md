The auto generated CRUD operations are documented via [Swag](https://github.com/swaggo/swag) which converts Go annotations to Swagger Documentation 2.0.  

## Getting started

Install swag by using:

```text
go install github.com/swaggo/swag/cmd/swag@latest
```

## Generates swagger documentation

Run swag init in the project's root folder which contains the `main.go` file. This will parse your comments and generate 
the required files (docs folder and docs/docs.go).

```text
swag init --parseDependency --parseInternal
```

Make sure to import into the `main.go` file, the generated `docs/docs.go` so that your specific configuration gets `init`'ed.

```go
import _ "example-module-name/docs"
```

Also adds the main annotations like:

```go
package main

import (
    "github.com/darksubmarine/torpedo-lib-go/app"
    "github.com/darksubmarine/torpedo-lib-go/conf"
    "log/slog"
    "os"

    _ "./docs"
)


// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

    // 1. App configuration
    config := conf.Map{}
    config.Add(3000, "port")
    
    // 2. Depdencies
    opts := app.ContainerOpts{Log: app.ContainerLogsOpts{W: os.Stdout, L: slog.LevelInfo}}
    
    application := app.NewApplicationContainer(opts)
    application.WithProvider(dependency.NewHttpServerProvider(config))
    application.WithProvider(dependency.NewHelloProvider())
    
    // 3. Run your application!
    application.Run()
}
```

Now we can run again `swag init` to refresh the generated documentation. 

```text
swag init --parseDependency --parseInternal
```

!!! tip "Parse internal files and dependencies"
    The swag command can be executed with `parseDependency` and `parseInternal` options:
    ```
    swag init --parseDependency --parseInternal

    --parseInternal                        Parse go files in internal packages, disabled by default
    --parseDependency, --pd                Parse go files inside dependency folder, disabled by default
    ```

After using `swag init` to generate Swagger 2.0 docs, lets go to add the documentation endpoint to the HTTP Server.
Import the following packages within the HTTP server provider:

```go
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files" // swagger embed files
```

And register the endpoint:

```go
ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

For instance:

```go hl_lines="12 13 47"
package dependency

import (
    "github.com/darksubmarine/torpedo-lib-go/app"
    "github.com/darksubmarine/torpedo-lib-go/conf"
    "github.com/darksubmarine/torpedo-lib-go/log"
    "github.com/gin-gonic/gin"

    "fmt"
    "net/http"

    ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
    swaggerFiles "github.com/swaggo/files" // swagger embed files
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


	p.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
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

```

Finally, run your app, and browse to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)


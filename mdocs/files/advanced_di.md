
!!! info "Application Container"
    For further information please read the architecture section [Application Container](arch_application_container.html)


## Entity module

The first time that an entity code is generated an injection module is created at `dependency` package. 
This ensures the below entity construction steps: 
 
1. **Repository** set up.
2. **Service** set up.
3. **API** endpoints set up.

For instance, following the Booking Fly application example, the `trip` injection module looks like:

```go
package dependency

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	tripHTTP "github.com/darksubmarine/booking-fly/domain/entities/trip/inputs/http/gin"
	tripRepo "github.com/darksubmarine/booking-fly/domain/entities/trip/outputs/memory"
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"
)

type TripProvider struct {
	app.BaseProvider

	// trip service instance to be provided.
	service trip.IService `torpedo.di:"provide"`

	// trip repository instance to be provided.
	repo trip.IRepository `torpedo.di:"provide"`

	// logger instance provided by LoggerProvider.
	logger log.ILogger `torpedo.di:"bind"`

	// storageKey is the crypto key to encode encrypted fields at storage level.
	storageKey []byte `torpedo.di:"bind,name=STORAGE_KEY"`

	// apiV1 group to register endpoints
	apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"`

	// private fields initialized by constructor
	cfg conf.Map
}

func NewTripProvider(config conf.Map) *TripProvider {
	return &TripProvider{cfg: config}
}

// Provide provides instances.
func (p *TripProvider) Provide(c app.IContainer) error {

	// -- Repo (output) ---
	p.repo = tripRepo.NewMemoryRepository(p.storageKey) //(1)!

	// -- Service (business logic)
	p.service = trip.NewService(p.repo, p.logger) //(2)!

	// -- Controller (input) --
	controller := tripHTTP.NewInputGin(p.service, p.logger)
	controller.Register(p.apiV1) //(3)!

	return nil
}


```

1. Respository set up. Here you can create a memory repo or a MongoDB repo, etc.
2. Service set up. This step requires a previous repository instance. And is the service constructor.
3. API set up. All entity endpoints are set here.

## Use Case module

The first time that a use case code is generated an injection module is created at `dependency` package. This injection module per use case
ensures a safe place to build each use case instance which its required dependencies.

Following the `Booking Fly` example, the generated module looks like:

```go
package dependency

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	"github.com/darksubmarine/booking-fly/domain/entities/user"
	"github.com/darksubmarine/booking-fly/domain/use_cases/booking_fly"

	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/log"

	// Uncomment following lines if your use case contains http input.
	//BookingFlyHTTP "github.com/darksubmarine/booking-fly/domain/use_cases/booking_fly/inputs/http"
	//"github.com/gin-gonic/gin"
)

type UseCaseBookingFlyProvider struct {
	app.BaseProvider

	// useCaseBookingFly provides an booking_fly.UseCase instance.
	useCaseBookingFly *booking_fly.UseCase `torpedo.di:"provide"`

	// logger instance provided by LoggerProvider.
	logger log.ILogger `torpedo.di:"bind"`

	// userSrv instance of user service.
	userSrv user.IService `torpedo.di:"bind"`

	// tripSrv instance of trip service.
	tripSrv trip.IService `torpedo.di:"bind"`

	// Uncomment following lines if your use case contains http input.
	// api router group to add endpoints under /api prefix
	//apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"`
}

func NewUseCaseBookingFlyProvider() *UseCaseBookingFlyProvider {
	return &UseCaseBookingFlyProvider{}
}

// Provide provides the use case instance.
func (p *UseCaseBookingFlyProvider) Provide(c app.IContainer) error {
	p.useCaseBookingFly = booking_fly.NewUseCase(p.logger, p.userSrv, p.tripSrv)

	/*
	// Add your use case API endpoints
	p.apiV1.POST("/booking",
		BookingFlyHTTP.NewController(p.useCaseBookingFly, p.logger).BookingFlyEndpoint)
    */
	return nil
}

```

## Domain module

The domain module also is created since the beginning. This module is where all domain dependencies are bound and the Domain class
is created. This module ensures the right domain initialization with:

1. Domain Context
2. Entities references
3. Use cases references
4. Input Adapter connection with domain port

The next one is an example of it from the Booking Fly App:

```go
package dependency

import (
	"github.com/darksubmarine/booking-fly/domain"
	"github.com/darksubmarine/booking-fly/domain/inputs/http"

	"github.com/darksubmarine/booking-fly/domain/entities/trip"

	"github.com/darksubmarine/booking-fly/domain/entities/user"

	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"

	"github.com/darksubmarine/booking-fly/domain/use_cases/booking_fly"
)

type DomainProvider struct {
	app.BaseProvider

	// -- Providers --

	// domainCtx provide domain context instance
	domainCtx *domain.Context `torpedo.di:"provide"`

	// domainSrv provide domain service instance
	domainSrv domain.IDomainService `torpedo.di:"provide"`

	// -- Bind services --

	// trip service wired instance.
	tripSrv trip.IService `torpedo.di:"bind"`

	// user service wired instance.
	userSrv user.IService `torpedo.di:"bind"`

	// -- Bind use cases --
	
	// BookingFly use case wired instance.
	ucBookingFly *booking_fly.UseCase `torpedo.di:"bind"`

	// logger instance provided by LoggerProvider.
	logger log.ILogger `torpedo.di:"bind"`

	// apiV1 group to register endpoints
	apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"`

	// private fields initialized by constructor
	cfg conf.Map
}

func NewDomainProvider(config conf.Map) *DomainProvider {
	return &DomainProvider{cfg: config}
}

// Provide provides instances.
func (p *DomainProvider) Provide(c app.IContainer) error {

	p.domainCtx = domain.NewContext(
		p.tripSrv,
		p.userSrv,
	)

	p.domainSrv = domain.NewService(p.domainCtx /*, p.ucFoo, p.ucBar */)

	// API registration
	domainController := http.NewDomainController(p.domainSrv, p.logger)
	domainController.Register(p.apiV1)

	return nil
}


```

## Additional modules

New modules can be added following the application requirements. However, Torpedo adds out of the box some additional modules:

### Storage Key

The storage key module introduce a place to provide the storage key that will be used for those entities that have encrypted fields.
The provided module is so simple and fetch the key from the config yaml file. But due to this module is written only once, you can
modify this logic in order to fetch the secret key from a different datasource.

??? abstract "Storage Key module"
    ```go
    package dependency
    
    import (
        "github.com/darksubmarine/torpedo-lib-go/app"
        "github.com/darksubmarine/torpedo-lib-go/conf"
    )
    
    type StorageKeyProvider struct {
        app.BaseProvider
    
        key []byte `torpedo.di:"provide,name=STORAGE_KEY"`
    
        // private fields initialized by constructor
        cfg conf.Map
    }
    
    func NewStorageKeyProvider(config conf.Map) *StorageKeyProvider {
        return &StorageKeyProvider{cfg: config}
    }
    
    // Provide provides the storage key instance.
    func (p *StorageKeyProvider) Provide(c app.IContainer) error {
        p.key = []byte(p.cfg.FetchStringOrElse("the-key-has-to-be-32-bytes-long!","key")) //(1)!
    
        return nil
    }
    
    ```

    1. Setting the key from provided config map. The data source can be updated, maybe an env var? or a vault service?


### Logger

A logger module is provided based on top of the offical `log/slog` library. Torpedo supports all loggers that implements the interface
[log.ILogger](https://github.com/darksubmarine/torpedo-lib-go/blob/main/log/interface.go)

```go
package log

type ILogger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}
```

So, if you have your own logger or another logger that doesn't support this interface, a wrapper can be created and injected updating this module.

### HTTP Server

As input adapter Torpedo has been created on top of [Gin Gonic library](https://github.com/gin-gonic/gin). This dependency module
provides an instance of a default Gin web server. So, you can modify it as you need, maybe adding middlewares?

??? abstract "HTTP Server"
    ```go
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
    
    ```


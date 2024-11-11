An HTTP REST API as an input adapter in a hexagonal architecture serves as the interface for handling HTTP requests from clients. 
This adapter translates HTTP requests into calls to the core application's use cases and returns appropriate HTTP responses. 
[Torpedo provides a HTTP REST API built on top of Gin Gonic.](basic_entity_input_restapi.html#how-torpedo-implements-http-adapter) 


## Setting up the HTTP Server

Before binding the entity or use case CRUD endpoints an HTTP Server should be provided via dependency injection to have it available in all
providers.

In order to achieve it, please take a look at [Dependency Injection - HTTP server](advanced_di.html#http-server)

## Setting up the entity API

Once that we have an HTTP Server provider, is time to register our entity endpoints. What we need is a binding of the 
HTTP Server in our entity provider.

The example below illustrates how to register it to enable the [entity CRUD endpoints](basic_entity_input_restapi.html#api-endpoints)

```go hl_lines="5 29 49 50"
package dependency

import (
	"bitbucket.org/darksubmarine/machine/domain/entities/sensor"
	input "bitbucket.org/darksubmarine/machine/domain/entities/sensor/inputs/http/gin" //(1)!
	sensorRepo "bitbucket.org/darksubmarine/machine/domain/entities/sensor/outputs/memory"
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"
)

type SensorProvider struct {
	app.BaseProvider

	// sensor service instance to be provided.
	service sensor.IService `torpedo.di:"provide"`

	// sensor repository instance to be provided.
	repo sensor.IRepository `torpedo.di:"provide"`

	// logger instance provided by LoggerProvider.
	logger log.ILogger `torpedo.di:"bind"`

	// storageKey is the crypto key to encode encrypted fields at storage level.
	storageKey []byte `torpedo.di:"bind,name=STORAGE_KEY"`

	// apiV1 group to register endpoints
	apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"` //(2)!
	
	// private fields initialized by constructor
	cfg conf.Map
}

func NewSensorProvider(config conf.Map) *SensorProvider {
	return &SensorProvider{cfg: config}
}

// Provide provides instances.
func (p *SensorProvider) Provide(c app.IContainer) error {

	// -- Repo (output) ---
	p.repo = sensorRepo.NewMemoryRepository(p.storageKey)
	
	// -- Service (business logic)
	p.service = sensor.NewService(p.repo, p.logger)

	// -- Controller (input) --
	controller := input.NewInputGin(p.service, p.logger) //(3)!
	controller.Register(p.apiV1) //(4)!

	return nil
}

```

1. Import the entity input HTTP (gin) module. 
2. Bind a reference to the HTTP API server `*gin.RouterGroup`
3. Creates the controller instance with a service dependency useful to call service methods withing CRUD endpoints
4. Register the controller endpoints as part of the bound API v1.

## Setting up the use case API

Inside the use case directory the `input` folder will be there where you should create the DTO object and the Controller. 
Please follow the [use case quick start](quickstart_use_cases.html#adding-use-case-endpoint-to-rest-api) that explains how to add it.

Now with the previous files created, and following the Onboarding use case example, we can register it as part of the REST API within its own dependency module
located in `dependency/use_case_onboarding.go`.

```go hl_lines="6 9 25 36 37"
package dependency

import (
    "github.com/darksubmarine/blog/domain/entities/user"
    "github.com/darksubmarine/blog/domain/use_cases/onboarding"
    inputHTTP "github.com/darksubmarine/blogp/domain/use_cases/onboarding/inputs/http" //(1)!
    "github.com/darksubmarine/torpedo-lib-go/app"
    "github.com/darksubmarine/torpedo-lib-go/log"
    "github.com/gin-gonic/gin" //(2)!
)

type UseCaseOnboardingProvider struct {
    app.BaseProvider

    // useCaseOnboarding provides an onboarding instance.
    useCaseOnboarding *onboarding.UseCase `torpedo.di:"provide"`

    // logger instance provided by LoggerProvider.
    logger log.ILogger `torpedo.di:"bind"`

    // userSrv instance of user service.
    userSrv user.IService `torpedo.di:"bind"`

    // api router group to add endpoints under /api prefix
    apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"` //(3)!
}

func NewUseCaseOnboardingProvider() *UseCaseOnboardingProvider {
    return &UseCaseOnboardingProvider{}
}

// Provide provides the use case instance.
func (p *UseCaseOnboardingProvider) Provide(c app.IContainer) error {
    p.useCaseOnboarding = onboarding.NewUseCase(p.logger, p.userSrv)

	controller := inputHTTP.NewController(p.useCaseOnboarding, p.logger)
    p.apiV1.POST("/onboarding", controller.RegisterNewUser) //(4)!
    
    return nil
}

```

1. Import the use case http package.
2. Import the Gin Gonic dependency.
3. Bind with the REST API instance.
4. Adds the use case endpoint.


The dependencies are handled by the `Application Container`. The `torpedo fire` command will generated the basis for the application 
and more dependency modules can be added.

Torpedo will generate out of the box the following dependency modules:

| Module | Description |
|--------|-------------|
| logger | Provides a log instance based on `log/slog`|
| http_server | Provides a HTTP server based on Gin |
| storage_key | Provides a 32 bytes key required for encryption at storage layer |

The file `dependency/init.go` contains all the injected modules required for the application and it is the place to add more custom modules.

Additionally, per each defined entity and use case an injection module will be created where the repository settings and others can be configured.
Following the Booking Fly app example, the auto generated `init.go` file should look like:

```go
package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"log/slog"
	"os"
)

func NewAppContainer(cfg conf.Map, environment string) app.IApp {

	// Options to initialize the internal application container logger.
	opts := app.ContainerOpts{Log: app.ContainerLogsOpts{W: os.Stdout, L: slog.LevelInfo}}

	// returns the app.IApp instance.
	return app.NewContainer(opts).
		WithProvider(NewLoggerProvider(cfg.FetchSubMapP("logger"))).
		WithProvider(NewStorageKeyProvider(cfg.FetchSubMapP("domain", "storage"))).
		WithProvider(NewHttpServerProvider(cfg.FetchSubMapP("http"))).
		
		// entity providers
		WithProvider(NewTripProvider(cfg.FetchSubMapP("domain"))).
		WithProvider(NewUserProvider(cfg.FetchSubMapP("domain"))).
		
		// use cases providers
		WithProvider(NewUseCaseBookingFlyProvider())

	/* here you can add all needed providers */
}

```

For further information please read the section [Application Container](/arch_application_container.html)

# Application Container

The `ApplicationContainer` handles the app life cycle letting you provide dependencies applying
the control of inversion principle.

The dependencies have access to the application life cycle at:

- **Start**: Once that all dependencies are registered, the app is executed by `.Run()` method, here is when the app is set to `start` status and all dependencies `OnStart` hooks are executed.
- **Stop**: This is the last app status, when a signal termination is received the dependencies `OnStop` hooks are executed.

Please check the [example](https://github.com/darksubmarine/torpedo-lib-go/tree/main/app/_example) application.

## Provider

The provider structs are where the instances to provide are created and wired. Each provider have two different fields
annotated with Go tags in order to define its functionality. The tags could be:

### `torpedo.di:"provide"`
The field annotated with this tag is marked as provider to be bound in other provider which requires its instance.
Your provider should implement the method `Provide()` where all fields tagged as `provide` **MUST** be initialized.

### `torpedo.di:"bind"`
The field annotated with bind tag will have injected automatically the provided instance of its data type.

> When you need to provide more than one instance of the same data type the provide tag could contain an attribute name.
```go
type LoggerProvider struct {
	app.BaseProvider

	// instance to be provided
	defaultInstance *slog.Logger `torpedo.di:"provide"`
	instance        *slog.Logger `torpedo.di:"provide,name=LOGGER"`

}
```

The providers must implement the interface `IProvider` but for simplicity can extend the base struct named as `app.BaseProvider` which
has implemented all the required methods as no operation mode. So, the developer only needs to overwrite at least the `Provide()` method.

### Provide Singleton vs Builder function

The previous `LoggerProvider` example illustrates how to create a Singleton instance of `slog.Logger`. So, each Provider that contains a `bind`
of `*slog.Logger` will have injected the same logger instance, which for a logger could be ok, but what happens if we would like to inject
different instances per bound dependency. Here is when the Builder function is useful. This function will be called in order to build a new instance
of the required object. The previous example can be modified as:

```go
type LoggerProvider struct {
	app.BaseProvider

	// singletonInstance is a pointer to a singleton instance of slog.Logger
	singletonInstance *slog.Logger `torpedo.di:"provide"`
	
	// builderInstance Instance builder function that returns a pointer to a new instance of slog.Logger 
	builderInstance func(c app.IContainer) *slog.Logger `torpedo.di:"provide,name=LOGGER"`

}

// Provide this method sets the declared providers.
func (p *LoggerProvider) Provide(c app.IContainer) error {
	
    p.singletonInstance = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

    p.builderInstance = func(c app.IContainer) *slog.Logger {
	    return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))	
    }
	
    return nil
}
```



For instance a good example could be a
HTTP Server:

```go
package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"log/slog"

	"fmt"
	"io"
	"net/http"
)

type HttpServerProvider struct {
	app.BaseProvider
	cfg conf.Map

	// instance to be provided
	mux *http.ServeMux `torpedo.di:"provide,name=HTTPSERVER"`
}

func NewHttpServerProvider(config conf.Map) *HttpServerProvider {
	return &HttpServerProvider{cfg: config}
}

func (p *HttpServerProvider) Provide(c app.IContainer) error {
	p.mux = http.NewServeMux()
	p.mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "pong!")
	})

	return nil
}

func (p *HttpServerProvider) OnStart() func() error {
	return func() error {
		go func() {
			port := p.cfg.FetchIntP("port")
			if err := http.ListenAndServe(fmt.Sprintf(":%d", port), p.mux); err != nil {
				panic(fmt.Sprintf("error starting HTTP server at %d with error %s", port, err))
			}
		}()

		return nil
	}
}

```

Once that we have the HTTP server provider we could create new provider to add an endpoint:

```go
package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"io"

	"net/http"
)

type HelloProvider struct {
	app.BaseProvider
	
	mux    *http.ServeMux `torpedo.di:"bind,name=HTTPSERVER"`
}

func NewHelloProvider() *HelloProvider { return new(HelloProvider) }

func (p *HelloProvider) Provide(c app.IContainer) error {
	p.mux.HandleFunc("/hello", p.sayHello)

	// in this case we won't be providing anything...
	// only using the Provide method to inject a hello controller
	// as part of the mux server
	return nil
}

func (p *HelloProvider) sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}
```

## Running the application container

The main application must be initialized with an instance of `IApp` like:

```go
package main

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"log/slog"
	"os"
)

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

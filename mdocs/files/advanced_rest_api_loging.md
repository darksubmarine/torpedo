The logging is one of the most important practice that developers follow to have available information about any possible issue or application behaviour.

In Go there are a lot of available loggers with different APIs and implementation patterns. Because of that, Torpedo defines a simple interface
that match with the official package `log/slog` and can be implemented by anyone or even any logger can be adapted to this interface via a wrapper struct.

```go
type ILogger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}
```

## Logger Provider

Torpedo creates a logger provider out-of-the-box injecting an instance of `log.ILogger` that can be bound fron other providers.

```go
package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"log/slog"
	"os"
)

type LoggerProvider struct {
	app.BaseProvider

	// logger instance to be provided
	logger log.ILogger `torpedo.di:"provide"`

	// private fields initialized by constructor
	cfg conf.Map
}

func NewLoggerProvider(config conf.Map) *LoggerProvider {
	return &LoggerProvider{cfg: config}
}

// Provide provides the logger instance.
func (p *LoggerProvider) Provide(c app.IContainer) error {
	// read log level from configuration. Info will be by default if not found.
	sLvl := conf.SlogLevel(p.cfg.FetchStringOrElse("info", "level"))

	// creating log instance as singleton.
	p.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: sLvl}))

	return nil
}

```
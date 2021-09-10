package rest

import (
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/transport/httpcore"
	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

type Handler interface {
	API() *fury.Application
	RouteURLs(app *fury.Application)
}

// API constructs an http.Handler with all application routes defined.
func API(handler Handler) *fury.Application {
	// disabled StacktraceOnError and WithCaller
	// because both by default are not accurate
	// core-libs is responsible for put both in the logs
	// A custom error handler is handled by core-libs as well
	// to have useful stacktraces in new relic
	app, err := fury.NewWebApplication(
		fury.WithLogOptions(
			log.StacktraceOnError(false),
			log.WithCaller(false),
		),
		fury.WithErrorHandler(httpcore.ErrorHandler),
		fury.WithErrorEncoder(httpcore.ErrorEncoder),
	)
	if err != nil {
		panic(err.Error())
	}

	handler.RouteURLs(app)
	return app
}

package rest

import (
	"github.com/diegodileoML/practice_CDB/pkg/rest"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/transport/httpcore"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func (h *handler) RouteURLs(app *fury.Application) {
	group := app.Router.Group("/api/practice/users")

		group.Post("/", h.Store, httpcore.Middle(app))
		group.Get("/{id}",h.GetByID,httpcore.Middle(app))
		group.Put("/{id}",h.Update,httpcore.Middle(app))
		group.Delete("/{id}",h.Delete,httpcore.Middle(app))


}

// API constructs an http.Handler with all application routes defined.
func (h *handler) API() *fury.Application {
	return rest.API(h)
}


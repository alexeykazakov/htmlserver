package server

import (
	"github.com/alexeykazakov/htmlserver/pkg/assets"

	"github.com/gin-contrib/static"
	errs "github.com/pkg/errors"
)

func (srv *Server) SetupRoutes() error {
	var err error
	srv.routesSetup.Do(func() {
		// Create the route for static content, served from /
		var staticHandler static.ServeFileSystem
		staticHandler, err = assets.ServeEmbedContent()
		if err != nil {
			err = errs.Wrap(err, "unable to setup route to serve static content")
		}
		srv.router.Use(static.Serve("/", staticHandler))
	})
	return err
}

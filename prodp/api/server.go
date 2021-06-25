package api

import (
	"net/http"
	"prodp/data"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	Repository *data.Client
}

type Options struct {
	Production bool
}

func (s Server) Serve(options Options) {
	e := echo.New()

	paths := getHandlers()

	for _, handler := range paths {
		handleFunc := handler.handleFunc
		switch handler.method {
		case http.MethodGet:
			e.GET(handler.path, func(c echo.Context) error { return handleFunc(c, s.Repository) })
		case http.MethodPost:
			e.POST(handler.path, func(c echo.Context) error { return handleFunc(c, s.Repository) })
		case http.MethodPut:
			e.PUT(handler.path, func(c echo.Context) error { return handleFunc(c, s.Repository) })
		}
	}

	if options.Production {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("prodp.yifansong.ca")
		e.Pre(middleware.HTTPSRedirect())
		go func() { e.Logger.Fatal(e.Start(":80")) }()
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		e.Logger.Fatal(e.StartTLS(":4433", "cert.pem", "key.pem"))
	}
}

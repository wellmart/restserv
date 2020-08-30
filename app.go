package restserv

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

// App represents a App http server.
type App struct {
	http.Handler
	route           *route
	notFoundHandler *Handler
	middlewares     []*Middleware
}

// Middleware represents a Middleware in the http server.
type Middleware func(c *Context, r *http.Request) bool

// New creates a new app http server.
func New() *App {
	return &App{route: &route{}}
}

// Use append a new middleware to server.
func (a *App) Use(middleware Middleware) {
	a.middlewares = append(a.middlewares, &middleware)
}

// Delete maps the delete http method to handler.
func (a *App) Delete(path string, handler Handler) {
	a.route.add(http.MethodDelete, path, handler)
}

// Get maps the get http method to handler.
func (a *App) Get(path string, handler Handler) {
	a.route.add(http.MethodGet, path, handler)
}

// Patch maps the patch http method to handler.
func (a *App) Patch(path string, handler Handler) {
	a.route.add(http.MethodPatch, path, handler)
}

// Post maps the post http method to handler.
func (a *App) Post(path string, handler Handler) {
	a.route.add(http.MethodPost, path, handler)
}

// Put maps the put http method to handler.
func (a *App) Put(path string, handler Handler) {
	a.route.add(http.MethodPut, path, handler)
}

// NotFound maps a generic not found to handler.
func (a *App) NotFound(handler Handler) {
	a.notFoundHandler = &handler
}

// ListenAndServe starts the rest API server.
func (a *App) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, a)
}

// RunFcgi starts the rest API server in Fcgi.
func (a *App) RunFcgi(addr string) (e error) {
	var listener net.Listener

	if addr[0] == '/' {
		os.Remove(addr)
		listener, e = net.Listen("unix", addr)
	} else {
		listener, e = net.Listen("tcp", addr)
	}

	if e == nil {
		e = fcgi.Serve(listener, a)
	}

	return
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	handler, routeFound := a.route.match(r.Method, r.URL.Path, values)

	c := &Context{w, values}

	for _, middleware := range a.middlewares {
		if (*middleware)(c, r) {
			return
		}
	}

	if handler != nil {
		(*handler)(c)
	} else if routeFound {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else if a.notFoundHandler != nil {
		(*a.notFoundHandler)(c)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

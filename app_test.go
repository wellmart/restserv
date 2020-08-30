package restserv

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAllHandlers(t *testing.T) {
	app := New()

	app.Delete("/", func(c *Context) {
		c.OK()
	})

	app.Get("/", func(c *Context) {
		c.OK()
	})

	app.Patch("/", func(c *Context) {
		c.OK()
	})

	app.Post("/", func(c *Context) {
		c.OK()
	})

	app.Put("/", func(c *Context) {
		c.OK()
	})

	var w *httptest.ResponseRecorder
	var r *http.Request

	methods := []string{"DELETE", "GET", "PATCH", "POST", "PUT"}

	for _, method := range methods {
		r, _ = http.NewRequest(method, "/", nil)
		w = httptest.NewRecorder()

		if app.ServeHTTP(w, r); w.Code != http.StatusOK {
			t.Errorf("Failed with \"%s\" method", method)
			return
		}
	}
}

func TestNotFoundHandler(t *testing.T) {
	app := New()

	app.NotFound(func(c *Context) {
		c.NotFound("Not Found")
	})

	var w *httptest.ResponseRecorder
	var r *http.Request

	r, _ = http.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()

	if app.ServeHTTP(w, r); w.Code != http.StatusNotFound {
		t.Error("Failed not found handler")
	}
}

func TestInvalidMethodHandler(t *testing.T) {
	app := New()

	var w *httptest.ResponseRecorder
	var r *http.Request

	r, _ = http.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()

	if app.ServeHTTP(w, r); w.Code != http.StatusNotFound {
		t.Error("Failed with invalid GET method")
		return
	}
}

func TestMethodNotAllowedHandler(t *testing.T) {
	app := New()

	app.Get("/", func(c *Context) {})

	var w *httptest.ResponseRecorder
	var r *http.Request

	r, _ = http.NewRequest("POST", "/", nil)
	w = httptest.NewRecorder()

	if app.ServeHTTP(w, r); w.Code != http.StatusMethodNotAllowed {
		t.Error("Failed with not allowed method")
		return
	}
}

func TestRun(t *testing.T) {
	go func() {
		app := New()

		if e := app.ListenAndServe(":9999"); e != nil {
			t.Error(e.Error())
		}
	}()

	time.Sleep(500 * time.Millisecond)
}

func TestRunFcgiUnix(t *testing.T) {
	go func() {
		app := New()

		if e := app.RunFcgi("/var/tmp/restserv.sock"); e != nil {
			t.Error(e.Error())
		}
	}()

	time.Sleep(250 * time.Millisecond)
}

func TestRunFcgiTcp(t *testing.T) {
	go func() {
		app := New()

		if e := app.RunFcgi(":9998"); e != nil {
			t.Error(e.Error())
		}
	}()

	time.Sleep(250 * time.Millisecond)
}

func TestRunInvalidFcgi(t *testing.T) {
	go func() {
		app := New()

		if e := app.RunFcgi(":666999"); e == nil {
			t.Error("Invalid bind addr")
		}
	}()

	time.Sleep(250 * time.Millisecond)
}

func TestMiddleware(t *testing.T) {
	app := New()

	app.Use(func(c *Context, r *http.Request) bool {
		c.OK()
		return true
	})

	var w *httptest.ResponseRecorder
	var r *http.Request

	r, _ = http.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()

	if app.ServeHTTP(w, r); w.Code != http.StatusOK {
		t.Error("Failed middleware")
	}
}

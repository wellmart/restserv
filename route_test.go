package restserv

import (
	"net/url"
	"testing"
)

func TestSimpleRoute(t *testing.T) {
	route := &route{}

	route.add("?", "/", func(c *Context) {})

	if handler, _ := route.match("?", "/", nil); handler == nil {
		t.Error("Invalid route")
	}
}

func TestSimpleRouteWithoutLastBar(t *testing.T) {
	route := &route{}

	route.add("?", "/test/", func(c *Context) {})

	if handler, _ := route.match("?", "/test", nil); handler == nil {
		t.Error("Invalid route")
	}
}

func TestSimpleSubRoute(t *testing.T) {
	route := &route{}

	route.add("?", "/path1/path2/path3/", func(c *Context) {})

	if handler, _ := route.match("?", "/path1/path2/path3/", nil); handler == nil {
		t.Error("Invalid route")
	}
}

func TestSimpleInvalidSubRoute(t *testing.T) {
	route := &route{}

	route.add("?", "/path1/path2/path3/", func(c *Context) {})

	if handler, routeFound := route.match("?", "/path1/path2/", nil); handler != nil || routeFound {
		t.Error("Invalid route")
	}
}

func TestSimpleRouteFound(t *testing.T) {
	route := &route{}

	route.add("GET", "/", func(c *Context) {})

	if handler, routeFound := route.match("POST", "/", nil); handler != nil || !routeFound {
		t.Error("Invalid route")
	}
}

func TestOneParamRoute(t *testing.T) {
	route := &route{}
	params := url.Values{}

	route.add("?", "/path1/path2/:param/path3/", func(c *Context) {})

	if handler, _ := route.match("?", "/path1/path2/1/path3/", params); handler == nil || params == nil || params.Get("param") != "1" {
		t.Error("Invalid route")
	}
}

func TestTwoParamRoute(t *testing.T) {
	route := &route{}
	params := url.Values{}

	route.add("?", "/path1/path2/:param1/path3/:param2/", func(c *Context) {})

	if handler, _ := route.match("?", "/path1/path2/1/path3/2/", params); handler == nil || params == nil || params.Get("param1") != "1" || params.Get("param2") != "2" {
		t.Error("Invalid route")
	}
}

func TestSecondParamInvalidRoute(t *testing.T) {
	route := &route{}

	route.add("?", "/path1/path2/:param1/path3/:param2/", func(c *Context) {})

	if e := route.add("?", "/path1/path2/:param3/path4/", func(c *Context) {}); e == nil {
		t.Error("Invalid route")
	}
}

func TestSimpleMultipleRoutes(t *testing.T) {
	route := &route{}

	route.add("?", "/path1/path2/path3/", func(c *Context) {})
	route.add("?", "/path1/path2/path4/", func(c *Context) {})

	if handler, _ := route.match("?", "/path1/path2/path3/", nil); handler == nil {
		t.Error("Invalid route")
	}

	if handler, _ := route.match("?", "/path1/path2/path4/", nil); handler == nil {
		t.Error("Invalid route")
	}
}

func TestSimpleOneParamRoutes(t *testing.T) {
	route := &route{}
	params := url.Values{}

	route.add("?", "/path1/path2/:param/path3/", func(c *Context) {})
	route.add("?", "/path1/path2/:param/path4/", func(c *Context) {})

	if handler, _ := route.match("?", "/path1/path2/1/path3/", params); handler == nil || params == nil || params.Get("param") != "1" {
		t.Error("Invalid route")
	}

	if handler, _ := route.match("?", "/path1/path2/2/path4/", params); handler == nil || params == nil || params.Get("param") != "2" {
		t.Error("Invalid route")
	}
}

func TestSimpleTwoParamRoutes(t *testing.T) {
	route := &route{}
	params := url.Values{}

	route.add("?", "/path1/path2/:param1/path3/:param2/", func(c *Context) {})
	route.add("?", "/path1/path2/:param1/path4/", func(c *Context) {})

	if handler, _ := route.match("?", "/path1/path2/1/path3/2/", params); handler == nil || params == nil || params.Get("param1") != "1" || params.Get("param2") != "2" {
		t.Error("Invalid route")
	}

	if handler, _ := route.match("?", "/path1/path2/3/path4/", params); handler == nil || params == nil || params.Get("param1") != "3" {
		t.Error("Invalid route")
	}
}

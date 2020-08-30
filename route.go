package restserv

import (
	"fmt"
	"net/url"
	"strings"
)

// Handler represents a router handler.
type Handler func(c *Context)

type route struct {
	part      string
	parameter bool
	handlers  map[string]*Handler
	children  []*route
}

func (r *route) add(method, path string, handler Handler) (e error) {
	var child *route

	parts := strings.Split(path, "/")[1:]
	level := r

	for i, part := range parts {
		if len(part) == 0 && i != 0 {
			continue
		}

		if child = level.get(part); child == nil {
			child = &route{
				part:      part,
				parameter: len(part) > 0 && part[0] == ':',
				handlers:  make(map[string]*Handler)}

			level.children = append(level.children, child)
		}

		if child.parameter && child.part != part {
			e = fmt.Errorf("\"%s\" already is reserved, \"%s\" is not accepted", child.part, part)
			return
		}

		level = child
	}

	level.handlers[method] = &handler
	return
}

func (r *route) match(method, path string, values url.Values) (handler *Handler, routeFound bool) {
	var child *route

	parts := strings.Split(path, "/")[1:]
	level := r

	for i, part := range parts {
		if len(part) == 0 && i != 0 {
			continue
		}

		if child = level.get(part); child == nil {
			return
		}

		if child.parameter {
			values.Set(child.part[1:], part)
		}

		level = child
	}

	if handler = level.handlers[method]; len(level.handlers) > 0 && handler == nil {
		routeFound = true
	}

	return
}

func (r *route) get(part string) (route *route) {
	for _, child := range r.children {
		if child.part == part || child.parameter {
			route = child
			break
		}
	}

	return
}

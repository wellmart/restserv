package restserv

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Context represents a context of request handler.
type Context struct {
	http.ResponseWriter
	values url.Values
}

// GetValue returns a value of context.
func (c *Context) GetValue(key string) string {
	return c.values.Get(key)
}

// SetValue sets a value of context.
func (c *Context) SetValue(key string, value string) {
	c.values.Set(key, value)
}

// BadRequest writes a bad request response.
func (c *Context) BadRequest() {
	c.WriteHeader(http.StatusBadRequest)
}

// OK writes a OK response.
func (c *Context) OK() {
	c.WriteHeader(http.StatusOK)
}

// NotModified writes a not modified response.
func (c *Context) NotModified() {
	c.WriteHeader(http.StatusNotModified)
}

// Unauthorized writes a unauthorized response.
func (c *Context) Unauthorized(message string) {
	c.writeMessage(http.StatusUnauthorized, message)
}

// NotFound writes a not found response.
func (c *Context) NotFound(message string) {
	c.writeMessage(http.StatusNotFound, message)
}

// Unprocessable writes a unprocessable response.
func (c *Context) Unprocessable(message string) {
	c.writeMessage(http.StatusUnprocessableEntity, message)
}

// Error writes a error response.
func (c *Context) Error(e error) {
	c.writeMessage(http.StatusInternalServerError, e.Error())
}

// WriteObject writes a json object to response.
func (c *Context) WriteObject(object interface{}) {
	buffer, e := json.Marshal(object)

	if e != nil {
		c.writeMessage(http.StatusInternalServerError, e.Error())
		return
	}

	c.writeBuffer(http.StatusOK, buffer)
}

// WriteRawMessage writes a json raw message to response.
func (c *Context) WriteRawMessage(message json.RawMessage) {
	c.writeBuffer(http.StatusOK, message)
}

func (c *Context) writeBuffer(code int, buffer []byte) {
	c.Header().Set("Content-Type", "application/json")

	c.WriteHeader(code)
	c.Write(buffer)
}

func (c *Context) writeMessage(code int, message string) {
	buffer, _ := json.Marshal(message)
	c.writeBuffer(code, []byte("{\"message\":"+string(buffer)+"}"))
}

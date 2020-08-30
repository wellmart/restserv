package restserv

import (
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.BadRequest()

	if w.Code != http.StatusBadRequest {
		t.Error("Failed BadRequest")
	}
}

func TestOK(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.OK()

	if w.Code != http.StatusOK {
		t.Error("Failed OK")
	}
}

func TestNotModified(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.NotModified()

	if w.Code != http.StatusNotModified {
		t.Error("Failed NotModified")
	}
}

func TestUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.Unauthorized("Unauthorized")

	if w.Code != http.StatusUnauthorized {
		t.Error("Failed Unauthorized")
	}
}

func TestNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.NotFound("Not Found")

	if w.Code != http.StatusNotFound {
		t.Error("Failed NotFound")
	}
}

func TestUnprocessable(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.Unprocessable("Unprocessable")

	if w.Code != http.StatusUnprocessableEntity {
		t.Error("Failed Unprocessable")
	}
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.Error(errors.New("Fake Error"))

	if w.Code != http.StatusInternalServerError {
		t.Error("Failed Error")
	}
}

func TestAbort(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.writeMessage(http.StatusOK, "OK")

	if w.Code != http.StatusOK || w.Body.String() != "{\"message\":\"OK\"}" {
		t.Error("Failed abort")
	}
}

func TestWriteObject(t *testing.T) {
	type test struct {
		Test1 string
		Test2 int
	}

	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.WriteObject(&test{Test1: "Hello World", Test2: 100})

	if w.Code != http.StatusOK || w.Body.String() != "{\"Test1\":\"Hello World\",\"Test2\":100}" {
		t.Error("Failed WriteObject")
	}
}

func TestWriteRawMessage(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.WriteRawMessage([]byte("{\"Test1\":\"Hello World\",\"Test2\":100}"))

	if w.Code != http.StatusOK || w.Body.String() != "{\"Test1\":\"Hello World\",\"Test2\":100}" {
		t.Error("Failed WriteRawMessage")
	}
}

func TestInvalidWriteObject(t *testing.T) {
	w := httptest.NewRecorder()
	c := &Context{w, nil}

	c.WriteObject(math.Inf(1))

	if w.Code != http.StatusInternalServerError {
		t.Error("Failed Invalid WriteObject")
	}
}

func TestValues(t *testing.T) {
	values := url.Values{}
	c := &Context{nil, values}

	c.SetValue("test", "ok")

	if c.GetValue("test") != "ok" {
		t.Error("Failed Values")
	}
}

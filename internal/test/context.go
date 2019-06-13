package test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

// NewContext returns a new echo.Context, and *httptest.ResponseRecorder to be
// used for tests.
func NewContext(t *testing.T, method string, r io.Reader, ctype string) (echo.Context, *httptest.ResponseRecorder) {
	t.Helper()

	e := echo.New()
	req := httptest.NewRequest(method, "/", r)
	req.Header.Set(echo.HeaderContentType, ctype)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}
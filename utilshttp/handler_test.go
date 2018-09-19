package utilshttp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AlphaWong/log-entry/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestPostMsgSuccess(t *testing.T) {
	var required = require.New(t)
	jsonschema.Init()
	var r, _ = http.NewRequest(http.MethodPost, "", strings.NewReader(`
	{
    "message": "1", 
    "src_file": "2",
    "src_line": "3",
    "context": {
    	"a": "b",
    	"c": "bn"
    },
    "level": "debug", 
    "time": "2018-09-19T07:45:46.215081910Z"
}
	`))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LogHandler)
	handler.ServeHTTP(rr, r)

	required.Equal(http.StatusOK, rr.Code)
}

func TestHealthCheckSuccess(t *testing.T) {
	var required = require.New(t)
	jsonschema.Init()
	var r, _ = http.NewRequest(http.MethodGet, "/healthCheck", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, r)

	required.Equal(http.StatusOK, rr.Code)
}

func TestHealthCheckFailByInvalidMethod(t *testing.T) {
	var required = require.New(t)
	jsonschema.Init()
	var r, _ = http.NewRequest(http.MethodPut, "/healthCheck", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, r)

	required.Equal(http.StatusMethodNotAllowed, rr.Code)
}

func TestPostMsgFailByInvalidJSON(t *testing.T) {
	var required = require.New(t)
	jsonschema.Init()
	var r, _ = http.NewRequest(http.MethodPost, "", strings.NewReader(`
	{
    "message": "1", 
    "src_file": "2",
    "src_line": "3",
    "context": {
    	"a": "b",
    	"c": "bn"
    },
    "level": "debug",,,,,,,
    "time": "2018-09-19T07:45:46.215081910Z"
}
	`))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LogHandler)
	handler.ServeHTTP(rr, r)

	required.Equal(http.StatusBadRequest, rr.Code)
}

func TestPostMsgFailByInvalidSrcLineType(t *testing.T) {
	var required = require.New(t)
	jsonschema.Init()
	var r, _ = http.NewRequest(http.MethodPost, "", strings.NewReader(`
	{
    "message": "1", 
    "src_file": "2",
    "src_line": "x",
    "context": {
    	"a": "b",
    	"c": "bn"
    },
    "level": "debug",
    "time": "2018-09-19T07:45:46.215081910Z"
}
	`))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LogHandler)
	handler.ServeHTTP(rr, r)

	required.Equal(http.StatusBadRequest, rr.Code)
}

func TestPostMsgFailByInvalidHttpMethod(t *testing.T) {
	var required = require.New(t)
	jsonschema.Init()
	var r, _ = http.NewRequest(http.MethodPut, "", strings.NewReader(`
	{
    "message": "1", 
    "src_file": "2",
    "src_line": "x",
    "context": {
    	"a": "b",
    	"c": "bn"
    },
    "level": "debug",
    "time": "2018-09-19T07:45:46.215081910Z"
}
	`))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LogHandler)
	handler.ServeHTTP(rr, r)

	required.Equal(http.StatusMethodNotAllowed, rr.Code)
}

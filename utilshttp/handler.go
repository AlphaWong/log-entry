package utilshttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/AlphaWong/log-entry/jsonschema"
	"github.com/AlphaWong/log-entry/types"
	"github.com/AlphaWong/log-entry/utils"
	lalamove "github.com/lalamove-go/logs"
	"go.uber.org/zap"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "")
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Request-ID", utils.GetUuid())
	if r.Method != http.MethodPost {
		RenderError(w, http.StatusMethodNotAllowed, "")
		return
	}

	var b, _ = ioutil.ReadAll(r.Body)
	var bodyStr = string(b)
	r.Body = ioutil.NopCloser(bytes.NewReader(b))

	if err := jsonschema.IsValidRequest(bodyStr); err != nil {
		var dump, _ = httputil.DumpRequest(r, true)
		lalamove.Logger().Error("Cannot pass json schema", zap.String("request", string(dump)))
		RenderError(w, http.StatusBadRequest, err.Error())
		lalamove.Logger().Sync()
		return
	}

	var req types.Request
	var err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		var dump, _ = httputil.DumpRequest(r, true)
		lalamove.Logger().Error("invalid request", zap.String("request", string(dump)))
		RenderError(w, http.StatusBadRequest, "")
		lalamove.Logger().Sync()
		return
	}

	var logMsg, _ = json.Marshal(req)
	fmt.Println(string(logMsg))

	w.Write([]byte(""))
}
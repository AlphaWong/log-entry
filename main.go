package main

import (
	"log"
	"net/http"

	"github.com/AlphaWong/log-entry/jsonschema"
	"github.com/AlphaWong/log-entry/utils"
	"github.com/AlphaWong/log-entry/utilshttp"
	lalamove "github.com/lalamove-go/logs"
	"go.uber.org/zap"
)

func main() {
	jsonschema.Init()
	http.HandleFunc("/", utilshttp.LogHandler)
	http.HandleFunc("/health", utilshttp.HealthCheckHandler)

	lalamove.Logger().Info("Server on at", zap.String("PORT", utils.Port))
	log.Fatal(http.ListenAndServe(utils.Port, nil))
}

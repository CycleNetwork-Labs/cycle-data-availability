package metrics

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	//Endpoint the endpoint for exposing the metrics
	Endpoint = "/metrics"
)

func StartMetricsHttpServer(c Config) {
	const ten = 10
	mux := http.NewServeMux()
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to create tcp listener for metrics: %v", err)
		return
	}
	mux.Handle(Endpoint, promhttp.Handler())

	metricsServer := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: ten * time.Second,
		ReadTimeout:       ten * time.Second,
	}
	log.Infof("metrics server listening on port %d", c.Port)
	go func() {
		if err := metricsServer.Serve(lis); err != nil {
			if err == http.ErrServerClosed {
				log.Warnf("http server for metrics stopped")
				return
			}
			log.Errorf("closed http connection for metrics server: %v", err)
			return
		}
	}()
}

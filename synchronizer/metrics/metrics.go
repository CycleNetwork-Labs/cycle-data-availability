package metrics

import (
	"github.com/0xPolygon/cdk-data-availability/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Prefix for the metrics of the block reconciliation package.
	Prefix = "synchronizer_"

	LatestSyncedBlockHeightName = Prefix + "latest_synced_block_height"
)

// Register the metrics for the block reconciliation package
func Register() {
	gauges := []prometheus.GaugeOpts{
		{
			Name: LatestSyncedBlockHeightName,
			Help: "[SYNCHRONIZER] latest synced block height",
		},
	}

	metrics.RegisterGauges(gauges...)
}

func LatestSyncedBlockHeight(blockNumber uint64) {
	metrics.GaugeSet(LatestSyncedBlockHeightName, float64(blockNumber))
}

package postgres

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

func NewMonitor(db *sql.DB, dbName string) prometheus.Collector {
	return collectors.NewDBStatsCollector(db, dbName)
}

package gormc

import (
	"github.com/zeromicro/go-zero/core/metric"
	"gorm.io/gorm"
	"time"
)

const metricsTimeKey = "metricsTime"
const metricsNamespace = "db"

type MetricsPlugin struct {
	queryTimeHistogram metric.HistogramVec
	errCount           metric.CounterVec
	sqlCount           metric.CounterVec
}

func (o *MetricsPlugin) Name() string {
	return "gorm_metrics"
}

func (o *MetricsPlugin) Initialize(db *gorm.DB) (err error) {
	o.sqlCount = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: metricsNamespace,
		Name:      "sql_count",
		Help:      "sql request counter",
		Labels:    []string{"table"},
	})

	o.queryTimeHistogram = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: metricsNamespace,
		Name:      "sql_duration_ms",
		Help:      "sql request duration ms",
		Labels:    []string{"table"},
		Buckets:   []float64{5, 10, 20, 50, 100, 200, 500, 1000, 2000},
	})

	o.errCount = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: metricsNamespace,
		Name:      "sql_err_count",
		Help:      "sql err counter",
		Labels:    []string{"table"},
	})

	if err = db.Callback().Query().Before("*").Register("query_metrics_before", o.Before); err != nil {
		return err
	}

	if err = db.Callback().Query().After("*").Register("query_metrics_after", o.After); err != nil {
		return err
	}

	if err = db.Callback().Create().Before("*").Register("create_metrics_before", o.Before); err != nil {
		return err
	}

	if err = db.Callback().Create().After("*").Register("create_metrics_after", o.After); err != nil {
		return err
	}

	if err = db.Callback().Update().Before("*").Register("update_metrics_before", o.Before); err != nil {
		return err
	}

	if err = db.Callback().Update().After("*").Register("update_metrics_after", o.After); err != nil {
		return err
	}

	if err = db.Callback().Delete().Before("*").Register("delete_metrics_before", o.Before); err != nil {
		return err
	}

	if err = db.Callback().Delete().After("*").Register("delete_metrics_after", o.After); err != nil {
		return err
	}

	return
}

func (o *MetricsPlugin) Before(db *gorm.DB) {
	now := time.Now()
	db.InstanceSet(metricsTimeKey, now)

}

func (o *MetricsPlugin) After(db *gorm.DB) {
	value, ok := db.InstanceGet(metricsTimeKey)
	if !ok {
		return
	}

	startTime := value.(time.Time)
	sqlTime := time.Since(startTime).Milliseconds()
	o.queryTimeHistogram.Observe(sqlTime, db.Statement.Table)

	o.sqlCount.Inc(db.Statement.Table)

	if db.Error != nil {
		o.errCount.Inc(db.Statement.Table)
	}
}

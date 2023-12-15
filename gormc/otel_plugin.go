package gormc

import (
	"github.com/zeromicro/go-zero/core/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	gormSpanKey = "gorm-span"
	gormCreate  = "gorm:create"
	gormUpdate  = "gorm:update"
	gormDelete  = "gorm:delete"
	gormQuery   = "gorm:query"
)

type OtelPlugin struct {
}

func (o OtelPlugin) Name() string {
	return "gorm-otel"
}

func (o OtelPlugin) Initialize(db *gorm.DB) (err error) {
	err = db.Callback().Create().Before("*").Register("otel_before_create", o.Before(gormCreate))
	if err != nil {
		return err
	}

	err = db.Callback().Create().After("*").Register("otel_after_create", o.After)
	if err != nil {
		return err
	}

	err = db.Callback().Update().Before("*").Register("otel_before_update", o.Before(gormUpdate))
	if err != nil {
		return err
	}

	err = db.Callback().Update().After("*").Register("otel_after_update", o.After)
	if err != nil {
		return err
	}

	err = db.Callback().Query().Before("*").Register("otel_before_query", o.Before(gormQuery))
	if err != nil {
		return err
	}

	err = db.Callback().Query().After("*").Register("otel_after_query", o.After)
	if err != nil {
		return err
	}

	err = db.Callback().Delete().Before("*").Register("otel_before_delete", o.Before(gormDelete))
	if err != nil {
		return err
	}

	err = db.Callback().Delete().After("*").Register("otel_after_delete", o.After)
	if err != nil {
		return err
	}

	return
}

func (o OtelPlugin) Before(spanName string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		ctx := db.Statement.Context
		tracer := trace.TracerFromContext(ctx)
		ctx, span := tracer.Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindClient))
		db.Statement.Context = ctx
		db.InstanceSet(gormSpanKey, span)
	}
}

func (o OtelPlugin) After(db *gorm.DB) {
	v, ok := db.InstanceGet(gormSpanKey)
	if !ok {
		return
	}

	span := v.(oteltrace.Span)
	defer span.End()

	if db.Statement.Error != nil {
		span.RecordError(db.Statement.Error)
	}

	span.SetAttributes(
		semconv.DBSQLTable(db.Statement.Table),
		semconv.DBStatement(db.Statement.SQL.String()),
	)
}

package gormc

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"
	"gorm.io/gorm"
)

func Transition(ctx context.Context, tx *gorm.DB, fc func(tx *gorm.DB) error) error {
	ctx, span := trace.TracerFromContext(ctx).Start(ctx, "transition")
	defer span.End()

	if err := tx.Transaction(fc); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

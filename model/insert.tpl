
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) error {
	{{if .withCache}}
    err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
        return db.Save(&data).Error
	}, m.getCacheKeys(data)...){{else}}db := m.conn
        err:= db.WithContext(ctx).Save(&data).Error{{end}}
	return err
}

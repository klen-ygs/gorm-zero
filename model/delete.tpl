
func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err != nil{
        if err == ErrNotFound {
            return nil
        }
		return err
	}
	 err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
        return db.Where("{{.originalPrimaryKey}} = @id", sql.Named("id", {{.lowerStartCamelPrimaryKey}})).Delete(&{{.upperStartCamelObject}}{}).Error
	}, m.getCacheKeys(data)...){{else}} db := m.conn
        err:= db.WithContext(ctx).Where("{{.originalPrimaryKey}} = @id", sql.Named("id", {{.lowerStartCamelPrimaryKey}})).Delete(&{{.upperStartCamelObject}}{}).Error
	{{end}}
	return err
}


func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryCtx(ctx, &resp, {{.cacheKeyVariable}}, func(conn *gorm.DB, v interface{}) error {
    		return conn.Model(&{{.upperStartCamelObject}}{}).Where("{{.originalPrimaryKey}} = @id", sql.Named("id", {{.lowerStartCamelPrimaryKey}})).First(&resp).Error
    	})
    if err == gormc.ErrNotFound {
    	return nil, err
    }
	return &resp, err
	{{else}}var resp {{.upperStartCamelObject}}
	err := m.conn.WithContext(ctx).Model(&{{.upperStartCamelObject}}{}).Where("{{.originalPrimaryKey}} = @id", sql.Named("id", {{.lowerStartCamelPrimaryKey}})).Take(&resp).Error
	if err == gormc.ErrNotFound {
        return nil, err
    }
	return &resp, err
	{{end}}
}

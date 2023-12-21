
func (m *default{{.upperStartCamelObject}}Model) FindOneBy{{.upperField}}(ctx context.Context, {{.in}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowIndexCtx(ctx, &resp, {{.cacheKeyVariable}}, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&{{.upperStartCamelObject}}{}).Where("{{.originalField}}", {{.lowerStartCamelField}}).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.{{.upperStartCamelPrimaryKey}}, nil
	}, m.queryPrimary)

	if err == gormc.ErrNotFound {
        return nil, err
    }
	return &resp, err
}{{else}}var resp {{.upperStartCamelObject}}
	err := m.conn.WithContext(ctx).Model(&{{.upperStartCamelObject}}{}).Where("{{.originalField}}", {{.lowerStartCamelField}}).Take(&resp).Error
	if err == gormc.ErrNotFound {
        return nil, err
    }
	return &resp, err
}{{end}}

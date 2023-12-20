{{ $space := " " -}}
{{ .name }}{{ $space -}}
{{if eq .name "DeletedAt" -}} {{- /* gorm 逻辑删除 */ -}}
gorm.DeletedAt{{ $space }}
{{- else -}}
{{ .type }}{{ $space }}
{{- end -}}
{{.tag}} {{if .hasComment}}// {{.comment}}{{end}}
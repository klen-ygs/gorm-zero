import (
	"context"
	"database/sql"
	{{if .time}}"time"{{end}}

    . "github.com/klen-ygs/gorm-zero/gormc/sql"

	"gorm.io/gorm"
)

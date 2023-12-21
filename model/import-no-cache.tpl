import (
	"context"
	"database/sql"
	{{if .time}}"time"{{end}}

    . "github.com/klen-ygs/gorm-zero/gormc/sql"
    "github.com/klen-ygs/gorm-zero/gormc"
	"gorm.io/gorm"
)

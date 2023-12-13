import (
	"context"
	"fmt"
	{{if .time}}"time"{{end}}
	"database/sql"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

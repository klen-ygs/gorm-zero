import (
	"context"
	"database/sql"
	"strings"
	"time"

    "github.com/klen-ygs/gorm-zero/gormc"
    . "github.com/klen-ygs/gorm-zero/gormc/sql"
	"gorm.io/gorm"
)

// avoid unused err
var _ = time.Second

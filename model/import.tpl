import (
	"context"
	"fmt"
	"time"
	"database/sql"

	. "github.com/klen-ygs/gorm-zero/gormc/sql"
	"github.com/klen-ygs/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

// avoid unused err
var _ = time.Second
package infrastructure

import (
	"lusionhouse-server/app/infrastructure/database/mysql"
	"lusionhouse-server/app/infrastructure/database/redis"

	"github.com/Lxxxxt/xutils"
)

func Startup() error {
	return xutils.Must(mysql.Startup, redis.Startup)
}

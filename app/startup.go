package app

import (
	"lusionhouse-server/app/infrastructure"

	"github.com/Lxxxxt/xutils"
)

func Startup() error {
	return xutils.Must(infrastructure.Startup)
}

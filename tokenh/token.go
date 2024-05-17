package tokenh

import (
	"github.com/gogf/gf/v2/util/gconv"
	"senchow/common/md5h"
	"senchow/common/randh"
)

func New(id uint64, salt string) string {
	return md5h.New(gconv.String(id)+salt) + randh.Str(32)
}

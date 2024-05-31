package tokenh

import (
	"github.com/gogf/gf/v2/util/grand"
	"github.com/tyronechan294/go-comm/md5h"
	"github.com/tyronechan294/go-comm/randh"
)

func New(id string, salt string) string {
	return md5h.Enc(id+salt) + grand.Str(randh.AlphaNumeric, 32)
}

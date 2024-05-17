package valueh

import (
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"time"
)

func IsNil(value interface{}, traceSource ...bool) bool {
	b := g.IsNil(value, traceSource...)
	if !b {
		logrus.Errorln(value)
	}
	return b
}

func IsNotNil(value interface{}, traceSource ...bool) bool {
	b := !g.IsNil(value, traceSource...)
	if b {
		logrus.Errorln(value)
	}
	return b
}

func IsEmpty(value interface{}) bool {
	return g.IsEmpty(value)
}

func IsNotEmpty(value interface{}) bool {
	return !g.IsEmpty(value)
}

func IsNotEmptyFatal(value interface{}) {
	if !g.IsEmpty(value) {
		logrus.Fatalln(value)
	}
}

func IsPanic(value interface{}, traceSource ...bool) {
	if !IsNil(value, traceSource...) {
		panic(value)
	}
}

// 复制对象
func Copy(fromValue interface{}, toValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: uint64(0),
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					ui, ok := src.(uint64)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return gconv.String(ui), nil
				},
			},
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(time.Time)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return s.Format(time.RFC3339), nil
				},
			},
		},
	})

}

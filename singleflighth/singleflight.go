package singleflighth

import "golang.org/x/sync/singleflight"

var sg = &singleflight.Group{}

func Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
	return sg.Do(key, fn)
}

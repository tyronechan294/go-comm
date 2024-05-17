package responseh

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

type Body struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Response(w http.ResponseWriter, resp any, err error) {
	var body Body
	if err != nil {
		//causeErr := errors.Cause(err)
		convert := ErrConvert(err) //status.Convert(causeErr)
		body.Code = int(convert.Code())
		body.Msg = convert.Message()
	} else {
		body.Msg = "Success"
		body.Data = resp

	}
	httpx.WriteJson(w, http.StatusOK, body)
}

func ErrConvert(err error) *status.Status {
	causeErr := errors.Cause(err)
	convert := status.Convert(causeErr)
	return convert
}
